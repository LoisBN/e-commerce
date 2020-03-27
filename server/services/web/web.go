package web

import (
	"log"
	"net/http"

	"github.com/loisBN/e-commerce/server/helpers"
	"github.com/loisBN/e-commerce/server/services/web/webgrpc"
)

var helper = &helpers.Helper{}

// RunAPI : start the api
func RunAPI()  {
	config,err := helper.GetConfig("web")
	if err != nil {
		log.Println(err.Error())
		return
	}
	webstub := &webgrpc.Webstub{}
	http.HandleFunc("/auth/signup",webstub.SignUp)
	http.HandleFunc("/auth/signin",webstub.SignIn)
	http.HandleFunc("/auth/signout",webstub.SignOut)
	http.ListenAndServe(config.(map[string]interface{})["api_endpoint"].(string),nil)
}

func messageAuth(w http.ResponseWriter,re *http.Request) {
	
}