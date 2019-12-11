package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository represents mongo db base operations interface
type Repository interface {
	FindOneByID(objectID string) *mongo.SingleResult
	// Insert(entity string)
	// Update(entity string)
	// Delete(id string)
}

// MongoBase respresents the struct for mongo db base operations
type MongoBase struct {
	client         *mongo.Client
	DBName         string
	CollectionName string
}

//FindOneByID gets the result by ObjectId
func (repository *MongoBase) FindOneByID(objectID string) *mongo.SingleResult {

	collectionName := repository.CollectionName
	db := repository.DBName

	collection := repository.client.Database(db).Collection(collectionName)

	objID, _ := primitive.ObjectIDFromHex(objectID)
	doc := bson.D{primitive.E{Key: "_id", Value: objID}}

	var res = collection.FindOne(context.TODO(), doc)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil //record not found
		}
	}

	return res
}

//FindOnebyDocument gets the result by document. Returns nil for record not found
func (repository *MongoBase) FindOnebyDocument(doc bson.D) *mongo.SingleResult {

	collectionName := repository.CollectionName
	db := repository.DBName

	collection := repository.client.Database(db).Collection(collectionName)

	var res = collection.FindOne(context.TODO(), doc)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil //record not found
		}
	}

	return res
}

//GetCollection returns the mongoDb collection reference
func (repository *MongoBase) GetCollection() *mongo.Collection {
	return repository.client.Database(repository.DBName).Collection(repository.CollectionName)
}

//SetClient sets the mongoDb client reference
func (repository *MongoBase) SetClient(c *mongo.Client) {
	repository.client = c
}

//we may consider exporting mongoClient

//GetClient returns the mongoDb client reference
// func (repository *MongoBase) GetClient() *mongo.Client {
// 	return repository.client
// }
