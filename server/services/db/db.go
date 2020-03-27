package db

import (
	"net"

	"github.com/loisBN/e-commerce/server/helpers"
	"github.com/loisBN/e-commerce/server/services/db/dbgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var c *helpers.Helper = &helpers.Helper{}

// LaunchServer : launch the db Service (and the db grpc server)
func LaunchServer() {
	config,err := c.GetConfig("db")
	if err != nil {
		grpclog.Errorln(err.Error())
		return
	}
	lis,err := net.Listen("tcp",config.(map[string]interface{})["server"].(map[string]interface{})["grpc_port"].(string))
	if err != nil {
		grpclog.Errorln(err.Error())
		return
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()
	dbHandler,err := dbgrpc.NewDBHandler(config.(map[string]interface{})["database_layer"].(string),config.(map[string]interface{})["database_info_connection"].(string))
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