package auth

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/loisBN/e-commerce/server/helpers"
	"github.com/loisBN/e-commerce/server/services/auth/authgrpc"
)



// LaunchAuthService : launch the auth grpc server
func LaunchAuthService()  {
	fmt.Println("launched")
	helper := &helpers.Helper{}
	config,err := helper.GetConfig("auth")
	if err != nil {
		grpclog.Errorln(err.Error())
		return
	}
	fmt.Println(config)
	li,err := net.Listen("tcp",config.(map[string]interface{})["server"].(map[string]interface{})["grpc_port"].(string))
	if err != nil {
		grpclog.Errorln(err.Error())
		return
	}
	defer li.Close()
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()
	authHandler := authgrpc.NewHandler()
	authgrpc.RegisterTestServer(grpcServer,authHandler)
	if err := grpcServer.Serve(li); err != nil {
		grpclog.Errorln(err.Error())
		grpcServer.Stop()
	}
	
}