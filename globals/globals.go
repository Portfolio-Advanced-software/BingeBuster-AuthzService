package globals

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Db             *mongo.Client
	AuthzDb        *mongo.Collection
	MongoCtx       context.Context
	MongoDBUrl     string
	DbName         string
	CollectionName string
	RabbitMQUrl    string
)
