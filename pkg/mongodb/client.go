package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// New creates a new client for mongo db
//local connStr: mongodb://localhost:27017
func New(connStr string, timeout int, logQueries bool) (*mongo.Client, error) {

	//todo: set timeout
	//todo: prepare queryLogger handler
	// Set client options
	clientOptions := options.Client().ApplyURI(connStr)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}
