package db

import (
	"context"
	"fmt"

	utils "github.com/LockettPundt/GOserver/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBconnect ----> connects to DB
func DBconnect() (*mongo.Client, error) {
	dbName := utils.GoDotEnvVariable("DB_NAME")
	dbURI := utils.GoDotEnvVariable("DB_URI")
	clientOptions := options.Client().ApplyURI(dbURI)
	// connecting to the graphQL tutorial DB that I have set up already.
	client, err := mongo.Connect(context.TODO(), clientOptions)
	fmt.Printf("Connected to %v", dbName)
	return client, err
}
