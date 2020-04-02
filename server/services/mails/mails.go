package mails

import (
	"net"

	"github.com/loisBN/e-commerce/server/helpers"
	"github.com/loisBN/e-commerce/server/services/mails/mailgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var helper = &helpers.Helper{}

func LaunchMailService() {
	config,err := helper.GetConfig("mails")
	if err != nil {
		grpclog.Errorln(err.Error())
		return
	}
	lis,err := net.Listen("tcp",config.(map[string]interface{})["server"].(map[string]interface{})["grpc_port"].(string))
	if err != nil {
		grpclog.Errorln(err.Error())
		return
	}
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()
	mailHandler := mailgrpc.NewMailHandler()
	mailgrpc.RegisterMailServer(grpcServer,mailHandler)
	if err := grpcServer.Serve(lis);err != nil {
		grpclog.Errorln(err.Error())
		grpcServer.Stop()
		return
	}
}
