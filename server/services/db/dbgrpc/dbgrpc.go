package dbgrpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/grpclog"
)


type DBHandler struct {DB DatabaseLayer}

// NewDBHandler : Create a new DBHandler
func NewDBHandler(dbType uint8 ,connection string) (*DBHandler,error) {
	db,err := GetDatabaseLayer(dbType,connection)
	return &DBHandler{
		DB: db,
	},err
}

// Add : Add A profile
func (db *DBHandler) Add(ctx context.Context,req *AuthRequest) (*AuthResponse,error) {
	err := db.DB.Add(req)
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil,err
	}
	return &AuthResponse{
		Info: "transmission success",
	},nil
} 

func (db *DBHandler) Get(ctx context.Context,req *AuthRequest) (*Profile,error) {
	return nil,errors.New("method not implemented")
}