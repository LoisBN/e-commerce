package main

import (
	"context"
	"fmt"
	"time"

	//"google.golang.org/grpc/grpclog"
	//"google.golang.org/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/loisBN/e-commerce/server/services/auth/authgrpc"
	"github.com/loisBN/e-commerce/server/services/web"
)

func main() {
	fmt.Println("test")
	//auth.Test()
	async := make(chan int)
	go web.RunAPI()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
	 	conn,err := grpc.Dial(":7890",grpc.WithInsecure())
	 	if err != nil {
	 		grpclog.Errorln(err.Error())
	 		return
	 	}
		 client := authgrpc.NewTestClient(conn)
		 res,_ := client.Hello(context.Background(),&authgrpc.Req{
			 Content: "this is the request",
		 })
		 fmt.Print(res)
	}
	
	<-async
}