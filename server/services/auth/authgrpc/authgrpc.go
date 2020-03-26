package authgrpc

import (
	"context"

	"github.com/loisBN/e-commerce/server/services/db/dbgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Handler : handler gRPC calls
type Handler struct{}

// Hello : Say Hello
func (h *Handler) Hello(ctx context.Context,req *Req) (*Res,error) {
	conn,err := grpc.Dial(":4444",grpc.WithInsecure())
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil,err
	}
	client := dbgrpc.NewAuthDBClient(conn)
	var x interface{}
	x = &dbgrpc.AuthRequest{}
	res,err := client.Add(context.Background(),x.(*dbgrpc.AuthRequest))
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil,err
	}
	grpclog.Infoln(res.GetInfo())
	return &Res{
		Content: "hello toi version 10",
	},nil
}

// NewHandler : Create a new gRPC handler for incoming connections
func NewHandler() *Handler {
	return &Handler{}
}