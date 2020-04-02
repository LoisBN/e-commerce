package webgrpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/loisBN/e-commerce/server/helpers"
	"github.com/loisBN/e-commerce/server/services/auth/authgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Webstub : allow you to perform rpc according to a http Request
type Webstub struct {
	Initiator string
}

const (
	// Auth : choose the auth service
	Auth uint8 = iota
	// DB : choose the db service
	DB
)

var helper = &helpers.Helper{}

// FormatDate : format the date according to the location
func FormatDate(t time.Time, location string) string {
	return "not implemented"
}

func getClient(service uint8) (authgrpc.AuthClient, error) {
	config,err := helper.GetConfig("web")
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil,err
	}
	conn, err := grpc.Dial(config.(map[string]interface{})["stub"].(map[string]interface{})["grpc_connection"].(string), grpc.WithInsecure())
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil,err
	}
	switch service {
	case Auth:
		return authgrpc.NewAuthClient(conn), nil
	default:
		return nil, errors.New("this service is not implemented yet")
	}
}

// SignUp : Sign up the user
func (wstub *Webstub) SignUp(w http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		return
	}
	client, err := getClient(Auth)
	if err != nil {
		grpclog.Errorln("error here",err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p := &authgrpc.Profile{}
	err = json.NewDecoder(req.Body).Decode(p)
	if err != nil {
		grpclog.Errorln("cannot decode json from request :", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var r *authgrpc.Req = &authgrpc.Req{
		Profile: p,
		Date:"now",
	}
	bf := time.Now().Nanosecond()
	res, err := client.SignUp(context.Background(), r)
	af := time.Now().Nanosecond()-bf
	fmt.Println(af)
	if err != nil {
		grpclog.Errorln(err.Error())
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	fmt.Println(res)
}

func (wstub *Webstub) UpdateProfile(w http.ResponseWriter,req *http.Request) {
	if req.Method == "OPTIONS" {
		return
	}
	client, err := getClient(Auth)
	if err != nil {
		grpclog.Errorln(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p := &authgrpc.Profile{}
	err = json.NewDecoder(req.Body).Decode(p)
	if err != nil {
		grpclog.Errorln("cannot decode json from request :", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var r *authgrpc.Req = &authgrpc.Req{
		Profile: p,
		Date:"now",
	}
	res,err := client.Update(context.Background(),r)
	if err != nil {
		grpclog.Errorln(err.Error())
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	fmt.Println(res)
}

func (wstub *Webstub) SignIn(w http.ResponseWriter,req *http.Request) {
	if req.Method == "OPTIONS" {
		return
	}
	client,err := getClient(Auth)
	if err != nil {
		grpclog.Errorln(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	x := &authgrpc.Profile{}
	err = json.NewDecoder(req.Body).Decode(x)
	if err != nil {
		grpclog.Errorln("cannot decode json from request :", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r := &authgrpc.Req{
		Profile: x,
		Date: "now",
	}
	ctx,cancel := context.WithDeadline(context.Background(),time.Now().Add(2*time.Second))
	defer cancel()
	profile,err := client.SignIn(ctx,r)
	if err != nil {
		grpclog.Errorln(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(profile)
}

func (wstub *Webstub) RecoverPassword(w http.ResponseWriter,req *http.Request) {
	if req.Method == "OPTIONS" {
		return
	}
	client,err := getClient(Auth)
	if err != nil {
		grpclog.Errorln(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	x := &authgrpc.Profile{}
	err = json.NewDecoder(req.Body).Decode(x)
	if err != nil {
		grpclog.Errorln("cannot decode json from request :", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r := &authgrpc.Req{
		Profile: x,
		Date: "now",
	}
	res,err := client.Recover(context.Background(),r)
	if err != nil {
		grpclog.Errorln(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// SignOut : allow the user to sign out
func (wstub *Webstub) SignOut(w http.ResponseWriter,req *http.Request) {
	if req.Method == "OPTIONS" {
		return
	}
	client,err := getClient(Auth)
	if err != nil {
		grpclog.Errorln(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	x := &authgrpc.Profile{}
	err = json.NewDecoder(req.Body).Decode(x)
	if err != nil {
		grpclog.Errorln("cannot decode json from request :", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res,err := client.SignOut(context.Background(),&authgrpc.Req{
		Profile: x,
	})
	if err != nil {
		grpclog.Errorln(err.Error())
		return
	}
	json.NewEncoder(w).Encode(res)
}
