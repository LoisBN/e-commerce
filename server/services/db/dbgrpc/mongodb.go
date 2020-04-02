package dbgrpc

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/loisBN/e-commerce/server/helpers"
	"github.com/loisBN/e-commerce/server/services/mails/mailgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var helper = &helpers.Helper{}

// MongoDBLayer : Handle database interaction with MongoDB
type MongoDBLayer struct{ *mgo.Session }

// NewMongoDBConnection : Create a new MongoDBLayer
func NewMongoDBConnection(connection string) (*MongoDBLayer, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &MongoDBLayer{session}, err
}

func getFreshSession(s *MongoDBLayer) *mgo.Session {
	return s.Copy()
}

func getMailClient() (mailgrpc.MailClient, error) {
	config, err := helper.GetConfig("mails")
	if err != nil {
		grpclog.Error(err.Error())
		return nil, err
	}
	conn, err := grpc.Dial(config.(map[string]interface{})["server"].(map[string]interface{})["grpc_port"].(string), grpc.WithInsecure())
	if err != nil {
		grpclog.Error(err.Error())
		return nil, err
	}
	return mailgrpc.NewMailClient(conn), nil
}

// Add : create a new record of a supported format
func (db *MongoDBLayer) Add(request interface{}) error {
	session := getFreshSession(db)
	switch r := request.(type) {
	case *AuthRequest:
		return addProfile(r, session)
	default:
		return errors.New("this type of query is not supported yet")
	}
}

func addProfile(req *AuthRequest, session *mgo.Session) error {
	err := session.DB(req.GetLocation()).C(req.GetCollection()).Insert(struct {
		Surname           string
		Name              string
		Email             string
		Birthday          string
		Date              string
		Password          string
		AccountValidation string
		Admin             string
	}{
		Surname:           req.GetProfile().GetSurname(),
		Name:              req.GetProfile().GetName(),
		Email:             req.GetProfile().GetEmail(),
		Birthday:          req.GetProfile().GetBirthday(),
		Date:              req.GetProfile().GetDate(),
		Password:          req.GetProfile().GetPassword(),
		AccountValidation: req.GetProfile().GetAccountValidation(),
		Admin:             req.GetProfile().GetAdmin(),
	})
	if err != nil {
		log.Println("error is right here", err.Error())
		return err
	}
	return nil
}

// Get : Get a record of a supported format
func (db *MongoDBLayer) Get(req interface{}) (map[string]interface{}, error) {
	session := getFreshSession(db)
	switch r := req.(type) {
	case *AuthRequest:
		return getProfile(r, session)
	default:
		return nil, errors.New("this type of query is not supported yet")
	}
}

func getProfile(req *AuthRequest, session *mgo.Session) (map[string]interface{}, error) {
	var x map[string]interface{}
	err := session.DB(req.GetLocation()).C(req.GetCollection()).Find(bson.M{"email": req.GetProfile().GetEmail()}).One(&x)
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	return x, nil
}

// Find : find a record in the database and return true if he find it or false if not
func (db *MongoDBLayer) Find(req interface{}) (bool, error) {
	session := getFreshSession(db)
	switch r := req.(type) {
	case *AuthRequest:
		return findProfile(r, session), nil
	default:
		return false, errors.New("request type not handled yet by the database layer")
	}
}

func findProfile(r *AuthRequest, session *mgo.Session) bool {
	err := session.DB(r.GetLocation()).C(r.GetCollection()).Find(bson.M{"email": r.GetProfile().GetEmail()}).One(nil)
	if err != nil {
		return false
	}
	return true
}

// Update : update a record in the databse
func (db *MongoDBLayer) Update(req interface{}) error {
	session := getFreshSession(db)
	switch r := req.(type) {
	case *AuthRequest:
		return updateProfile(r, session)
	default:
		return errors.New("this type of request is not handled yet by the database layer")
	}
}

func updateProfile(req *AuthRequest, session *mgo.Session) error {
	var x map[string]string
	err := session.DB(req.GetLocation()).C(req.GetCollection()).Find(bson.M{"_id": bson.ObjectIdHex(req.GetProfile().GetXId())}).One(&x)
	if err != nil {
		grpclog.Errorln(err.Error())
		return err
	}
	bs, err := helper.EncryptPassword(req.GetProfile().GetPassword())
	if err != nil {
		grpclog.Errorln(err.Error())
		return err
	}
	z := make(map[string]interface{})
	z["surname"] = req.GetProfile().GetSurname()
	z["name"] = req.GetProfile().GetName()
	z["email"] = req.GetProfile().GetEmail()
	z["birthday"] = req.GetProfile().GetBirthday()
	z["password"] = string(bs)
	z["admin"] = req.GetProfile().GetAdmin()
	z["accountvalidation"] = req.GetProfile().GetAccountValidation()

	for i, val := range z {
		switch r := val.(type) {
		case string:
			if z[i].(string) != "" {
				err = session.DB(req.GetLocation()).C(req.GetCollection()).Update(bson.M{"_id": bson.ObjectIdHex(req.GetProfile().GetXId())}, bson.M{"$set": bson.M{i: r}})
			}
		case bool:
			if _, ok := z[i].(bool); ok {
				err = session.DB(req.GetLocation()).C(req.GetCollection()).Update(bson.M{"_id": bson.ObjectIdHex(req.GetProfile().GetXId())}, bson.M{"$set": bson.M{i: r}})
			}
		}
	}
	if err != nil {
		grpclog.Errorln(err.Error())
		return err
	}
	if x["email"] != req.GetProfile().GetEmail() {
		go func(req *AuthRequest) {
			client, err := getMailClient()
			if err != nil {
				grpclog.Error(err.Error())
				return
			}
			client.Send(context.Background(), &mailgrpc.Mailreq{
				To:      req.GetProfile().GetEmail(),
				Subject: mailgrpc.UPDATEVERIFY,
			})
			fmt.Println("email Sent")
		}(req)
	}
	return nil
}
