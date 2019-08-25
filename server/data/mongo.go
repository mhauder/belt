package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client, _ = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"), nil)
