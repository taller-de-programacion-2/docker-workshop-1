package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

var mongoClient *mongo.Client

func main() {
	// Get PORT from environment variable
	port := os.Getenv("SERVER_PORT")

	initMongoClient()

	// Define handlers
	mux := http.NewServeMux()

	// Service healthcheck
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("pong")
		writer.Write([]byte("it works!"))
	})

	// Database (mongo) healthcheck
	mux.HandleFunc("/test-db", func(writer http.ResponseWriter, request *http.Request) {
		// Check the connection
		err := mongoClient.Ping(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}
		writer.Write([]byte("Connected to MongoDB!"))
	})

	fmt.Printf("Listening on port %s... \n", port)
	http.ListenAndServe(port, mux)
}

func initMongoClient() {
	// Set client options
	mongoUrl := os.Getenv("MONGO_URL")
	clientOptions := options.Client().ApplyURI(mongoUrl)

	// Connect to MongoDB
	var err error
	mongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo client created successfully")
}
