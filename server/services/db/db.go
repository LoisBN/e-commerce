package db

import (
	"net"

	"github.com/loisBN/e-commerce/server/services/db/dbgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func LaunchServer() {
	lis,err := net.Listen("tcp",":4444")
	if err != nil {
		grpclog.Errorln(err.Error())
		return
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()
	dbHandler,err := dbgrpc.NewDBHandler(0,"mongodb://localhost")
	if err != nil {
		grpclog.Errorln(err.Error())
		return
	}
	dbgrpc.RegisterAuthDBServer(grpcServer,dbHandler)
	if err := grpcServer.Serve(lis);err != nil {
		grpclog.Errorln(err.Error())
		grpcServer.Stop()
	}
}