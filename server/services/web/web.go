package web

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/loisBN/e-commerce/server/services/auth/authgrpc"
)

// RunAPI : start the api
func RunAPI()  {
	http.HandleFunc("/",messageAuth)
	http.ListenAndServe(":5678",nil)
}

func messageAuth(w http.ResponseWriter,re *http.Request) {
	conn,err := grpc.Dial(":7890",grpc.WithInsecure())
	if err != nil {
		grpclog.Errorln(err.Error())
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	client := authgrpc.NewTestClient(conn)
	mess,_ := client.Hello(context.Background(),&authgrpc.Req{
		Content: "yo les mecs",
	})
	fmt.Println("")
	fmt.Println(mess)
}