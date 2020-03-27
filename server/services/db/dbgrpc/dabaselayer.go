package dbgrpc

import (
	"errors"
)

const (
	// MONGODB : connection to mongoDB
	MONGODB string = "MongoDB"
	// MYSQL : connection to MySQL
	MYSQL
)

// DatabaseLayer : Interface to interact with database
type DatabaseLayer interface {
	Add(interface{}) error
	Get(interface{}) (map[string]interface{},error)
	Find(interface{}) (bool,error)
}

// GetDatabaseLayer : Get a connection to the selected database
func GetDatabaseLayer(dbLayer string,connection string) (DatabaseLayer,error) {
	switch dbLayer {
	case MONGODB:
		return NewMongoDBConnection(connection)
	default:
		return nil,errors.New("This database layer is not supported")
	}
}

