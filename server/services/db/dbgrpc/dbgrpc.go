package dbgrpc

import (
	"context"

	"google.golang.org/grpc/grpclog"
	"gopkg.in/mgo.v2/bson"
)

type AuthDBHandler struct{ DB DatabaseLayer }

// NewDBHandler : Create a new DBHandler
func NewDBHandler(dbType string, connection string) (*AuthDBHandler, error) {
	db, err := GetDatabaseLayer(dbType, connection)
	return &AuthDBHandler{
		DB: db,
	}, err
}

// Add : Add a record to the database
func (db *AuthDBHandler) Add(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {
	err := db.DB.Add(req)
	if err != nil {
		grpclog.Errorln("error is from here", err.Error())
		return nil, err
	}
	return &AuthResponse{
		Info: "transmission success",
	}, nil
}

// Update : update a database record
func(db *AuthDBHandler) Update(ctx context.Context,req *AuthRequest) (*AuthResponse,error) {
	err := db.DB.Update(req)
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	return &AuthResponse{
		Info: "transmission success",
	}, nil
}

// Get : Get a user from the database
func (db *AuthDBHandler) Get(ctx context.Context, req *AuthRequest) (*Profile, error) {
	result, err := db.DB.Get(req)
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	//fmt.Println("HERE ID",result["_id"].(bson.ObjectId).Hex())
	return &Profile{
		XId: 	  result["_id"].(bson.ObjectId).Hex(),
		Surname:  result["surname"].(string),
		Name:     result["name"].(string),
		Email:    result["email"].(string),
		Birthday: result["birthday"].(string),
		Date:     result["date"].(string),
		Password: result["password"].(string),
		Admin: 	  result["admin"].(string),
		AccountValidation: result["accountvalidation"].(string),
	}, nil
}

// Find : find an user inside the database
func (db *AuthDBHandler) Find(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {
	found, err := db.DB.Find(req)
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	return &AuthResponse{
		Info:    "transmission success",
		Success: found,
	}, nil
}
