package dbgrpc

import (
	"errors"
)

const (
	// MONGODB : connection to mongoDB
	MONGODB uint8 = iota
	// MYSQL : connection to MySQL
	MYSQL
)

type DatabaseLayer interface {
	Add(interface{}) error
	Get(interface{}) error
}

// GetDatabaseLayer : Get a connection to the selected database
func GetDatabaseLayer(dbLayer uint8,connection string) (DatabaseLayer,error) {
	switch dbLayer {
	case MONGODB:
		return NewMongoDBConnection(connection)
	default:
		return nil,errors.New("This database layer is not supported")
	}
}

