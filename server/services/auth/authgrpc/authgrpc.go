package authgrpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/loisBN/e-commerce/server/helpers"
	"github.com/loisBN/e-commerce/server/services/db/dbgrpc"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Handler : handler gRPC calls
type Handler struct{}

var helper = &helpers.Helper{}

func getDBConnection() (dbgrpc.AuthDBClient, error) {
	config, err := helper.GetConfig("db")
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	conn, err := grpc.Dial(config.(map[string]interface{})["server"].(map[string]interface{})["grpc_port"].(string), grpc.WithInsecure())
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	client := dbgrpc.NewAuthDBClient(conn)
	return client, nil
}

func formatRequest(location string, collection string, req *Req) *dbgrpc.AuthRequest {
	return &dbgrpc.AuthRequest{
		Location:   location,
		Collection: collection,
		Profile: &dbgrpc.Profile{
			Surname:  req.GetProfile().GetSurname(),
			Name:     req.GetProfile().GetName(),
			Email:    req.GetProfile().GetEmail(),
			Birthday: req.GetProfile().GetBirthday(),
			Date:     helper.GetFormatedDate(time.Now(), "paris", "Mon, 02 Jan 2006 15:04:05 MST"),
			Password: req.GetProfile().GetPassword(),
		},
	}
}

// SignUp : Add an user to the database
func (h *Handler) SignUp(ctx context.Context, req *Req) (*Res, error) {
	client, err := getDBConnection()
	if err != nil {
		grpclog.Errorln(err.Error)
		return nil, err
	}
	x := formatRequest("e-commerce", "users", req)
	res, err := client.Find(context.Background(), x)
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	if res.GetSuccess() == false {
		redis := helper.GetRedisClient(0)
		uuid := uuid.NewV4()
		redis.Set(req.GetProfile().GetEmail(), uuid.String(), 24*time.Hour)
		bpass, err := helper.EncryptPassword(req.GetProfile().GetPassword())
		if err != nil {
			grpclog.Errorln(err.Error())
			return nil, err
		}
		x.Profile.Password = string(bpass)
		res, err := client.Add(context.Background(), x)
		if err != nil {
			grpclog.Errorln(err.Error())
			return nil, err
		}
		return &Res{
			Content: res.GetInfo(),
			Profile: &Profile{
				Surname:  x.GetProfile().GetSurname(),
				Name:     x.GetProfile().GetName(),
				Email:    x.GetProfile().GetEmail(),
				Birthday: x.GetProfile().GetBirthday(),
				Date:     x.GetProfile().GetDate(),
			},
		}, nil
	}
	grpclog.Infoln(res.GetInfo())
	return &Res{
		Content: fmt.Sprintf("email already in use"),
	}, nil
}

// SignIn : authentcate the user if he have an accound
func (h *Handler) SignIn(ctx context.Context, req *Req) (*Res, error) {
	client, err := getDBConnection()
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	x := formatRequest("e-commerce", "users", req)
	fmt.Println("HERE IS X :",x)
	profile, err := client.Get(context.Background(), x)
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	fmt.Println("HERE IS PROFILE :",profile)
	err = bcrypt.CompareHashAndPassword([]byte(profile.GetPassword()), []byte(req.GetProfile().GetPassword()))
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	token := uuid.NewV4().String()
	redis := helper.GetRedisClient(0)
	redis.Set(profile.GetEmail(), token, 24*time.Hour)
	return &Res{
		Content: token,
		Profile: &Profile{
			Surname:  profile.GetSurname(),
			Name:     profile.GetName(),
			Email:    profile.GetEmail(),
			Birthday: profile.GetBirthday(),
			Date:     profile.GetDate(),
		},
	}, nil
}

// SignOut : sign out the user
func (h *Handler) SignOut(ctx context.Context, req *Req) (*Res, error) {
	redis := helper.GetRedisClient(0)
	redis.Del(req.GetProfile().GetEmail())
	return &Res{
		Content: "signout success",
	},nil
}

func (h *Handler) Update(ctx context.Context, req *Req) (*Res, error) {
	return nil, errors.New("method Add not implemented")
}

func (h *Handler) Authenticate(ctx context.Context, req *Req) (*Res, error) {
	return nil, errors.New("method Add not implemented")
}

// NewHandler : Create a new gRPC handler for incoming connections
func NewHandler() *Handler {
	return &Handler{}
}
