package main

import (
	"fmt"

	//"google.golang.org/grpc/grpclog"
	//"google.golang.org/grpc"

	"github.com/loisBN/e-commerce/server/services/auth"
	"github.com/loisBN/e-commerce/server/services/db"
	"github.com/loisBN/e-commerce/server/services/web"
)

func main() {
	fmt.Println("test")
	//auth.Test()
	async := make(chan int)
	go web.RunAPI()
	go db.LaunchServer()
	go auth.LaunchAuthService()
	<-async
}