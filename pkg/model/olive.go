package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//Olive represents olive model
type Olive struct {
	ID      *primitive.ObjectID `bson:"_id"`
	Kind    string              `bson:"kind"`
	Country string              `bson:"country"`
}
