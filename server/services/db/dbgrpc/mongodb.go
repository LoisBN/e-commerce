package dbgrpc

import (
	"errors"
	"log"

	"google.golang.org/grpc/grpclog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
		Surname  string
		Name     string
		Email    string
		Birthday string
		Date     string
		Password string
	}{
		Surname:  req.GetProfile().GetSurname(),
		Name:     req.GetProfile().GetName(),
		Email:    req.GetProfile().GetEmail(),
		Birthday: req.GetProfile().GetBirthday(),
		Date:     req.GetProfile().GetDate(),
		Password: req.GetProfile().GetPassword(),
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

func (db *MongoDBLayer) Find(req interface{}) (bool, error) {
	session := getFreshSession(db)
	switch r := req.(type) {
	case *AuthRequest:
		return findProfile(r, session), nil
	default:
		return false, errors.New("request type not handled yet")
	}
}

func findProfile(r *AuthRequest, session *mgo.Session) bool {
	err := session.DB(r.GetLocation()).C(r.GetCollection()).Find(bson.M{"email": r.GetProfile().GetEmail()}).One(nil)
	if err != nil {
		return false
	}
	return true
}
