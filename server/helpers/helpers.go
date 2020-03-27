package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/grpclog"
)

// Helper : contain helper function for making code easier
type Helper struct{}

const (
	// Full date and time
	Full string = "2 Jan 2006 15:04:05"
)

// GetFormatedDate : format the provided time with the provided layout
func (h *Helper) GetFormatedDate(t time.Time,location string,layout string) string {
	return t.Local().Format(layout)
}

// GetRedisClient : Return an instance of a redis client
func (h *Helper) GetRedisClient(db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		DB: db,
		Addr: "localhost:6379",
		Password: "",
	})
}

func (h *Helper) EncryptPassword(password string) ([]byte,error) {
		bpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			grpclog.Errorln(err.Error())
			return nil, err
		}
		return bpass,nil
}

// GetConfig : get the config from config.json according to the provided service
func (h *Helper) GetConfig(service string) (interface{},error) {
	var x map[string]interface{}
	f,err := os.Open("./config/config.json")
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	err = json.NewDecoder(f).Decode(&x)
	fmt.Println(x)
	if err != nil {
		grpclog.Errorln(err.Error())
		return nil, err
	}
	switch service {
	case "auth":
		return x["services"].(map[string]interface{})["auth"],nil
	case "db":
		return x["services"].(map[string]interface{})["db"],nil
	case "web":
		return x["services"].(map[string]interface{})["web"],nil
	default:
		return nil,errors.New("this services is not provided by configuration")
	}
}