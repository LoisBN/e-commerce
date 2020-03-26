package dbgrpc

import (
	"errors"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

// MongoDBLayer : Handle database interaction with MongoDB
type MongoDBLayer struct{*mgo.Session}

// NewMongoDBConnection : Create a new MongoDBLayer
func NewMongoDBConnection(connection string) (*MongoDBLayer,error) {
	session,err := mgo.Dial(connection) 
	if err != nil {
		log.Println(err.Error())
		return nil,err
	}
	return &MongoDBLayer{session},err
}

func getFreshSession(s *MongoDBLayer) *mgo.Session {
	return s.Copy()
}

func (db *MongoDBLayer) Add(request interface{}) error {
	session := getFreshSession(db)
	switch r := request.(type) {
	case *AuthRequest:
		return addProfile(r,session)
	default:
		return errors.New("the error is located right here")
	}
}

func addProfile(req *AuthRequest,session *mgo.Session) error {
	err := session.DB(req.GetLocation()).C(req.GetCollection()).Insert(req.GetProfile())
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (db *MongoDBLayer) Get(r interface{}) error {
	fmt.Println("not implemented yet")
	return errors.New("method not implemented yet")
}