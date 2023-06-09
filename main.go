package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/config"
	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/globals"
	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/handlers"
	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/messaging"
	mongodb "github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/mongodb"
	authzpb "github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/proto"
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

	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port " + c.Port + "...")

	// Set listener to start server
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("Unable to listen on port %p: %v", lis.Addr(), err)
	}

	// Set options, here we can configure things like TLS support
	opts := []grpc.ServerOption{}
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)
	// Create HistoryService type
	srv := &handlers.AuthzServiceServer{}

	// Register the service with the server
	authzpb.RegisterAuthzServiceServer(s, srv)

	// Construct the MongoDB URL
	globals.MongoDBUrl = fmt.Sprintf("mongodb+srv://%s:%s@%s", c.MongoDBUser, c.MongoDBPwd, c.MongoDBCluster)

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")
	globals.Db = mongodb.ConnectToMongoDB(globals.MongoDBUrl)

	globals.DbName = c.MongoDBDb
	globals.CollectionName = c.MongoDBCollection

	// Bind our collection to our global variable for use in other methods
	globals.AuthzDb = globals.Db.Database(globals.DbName).Collection(globals.CollectionName)

	// Construct the RabbitMQ URL
	globals.RabbitMQUrl = fmt.Sprintf("amqps://%s:%s@rattlesnake.rmq.cloudamqp.com/%s", c.RabbitMQUser, c.RabbitMQPwd, c.RabbitMQUser)

	//Connect to RabbitMQ
	fmt.Println("Connecting to RabbitMQ...")
	conn, err := messaging.ConnectToRabbitMQ(globals.RabbitMQUrl)
	if err != nil {
		log.Fatalf("Can't connect to RabbitMQ: %s", err)
	}

	// Start listening for messages RabbitMQ
	go messaging.ConsumeMessage(conn, "authz_queue", messaging.HandleMessage)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port " + c.Port)

	// Right way to stop the server using a SHUTDOWN HOOK
	// Create a channel to receive OS signals
	cs := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(cs, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-cs

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	lis.Close()
	fmt.Println("Closing MongoDB connection")
	globals.Db.Disconnect(globals.MongoCtx)
	fmt.Println("Done.")

}
