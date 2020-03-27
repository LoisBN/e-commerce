package main

import (
	"github.com/loisBN/e-commerce/server/services/auth"
	"github.com/loisBN/e-commerce/server/services/db"
	"github.com/loisBN/e-commerce/server/services/web"
)

func main()  {
	async := make(chan int)
	go web.RunAPI()
	go db.LaunchServer()
	go auth.LaunchAuthService()
	<-async
}