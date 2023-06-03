package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/config"
	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/handlers"
	mongodb "github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/mongodb"
	authzpb "github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/proto"
	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var db *mongo.Client
var authdb *mongo.Collection
var mongoCtx context.Context

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	// Construct the MongoDB URL
	mongodbURL := fmt.Sprintf("mongodb+srv://%s:%s@%s", c.MongoDBUser, c.MongoDBPwd, c.MongoDBCluster)

	// Initialize MongoDB client
	fmt.Println("Connecting to MongoDB...")
	db = mongodb.ConnectToMongoDB(mongodbURL)

	// Bind our collection to our global variable for use in other methods
	authdb = db.Database(c.MongoDBDb).Collection(c.MongoDBCollection)

	s := handlers.Server{
		DB:  authdb,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	authzpb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
