package auth

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/loisBN/e-commerce/server/services/auth/authgrpc"
)

func LaunchAuthService()  {
	li,err := net.Listen("tcp",":7890")
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