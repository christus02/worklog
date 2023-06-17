package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Load the ENV variables of this project
	loadEnv("worklog.env", true)

	pingDB()

}

func loadEnv(fileName string, verbose bool) {
	err := godotenv.Load("worklog.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envKeys := []string{"DB_USERNAME", "DB_PASSWORD", "MONGODB_ENDPOINT_URL", "MONGODB_CLUSTER_NAME", "MONGODB_SERVER_PREFIX"}

	if verbose {
		for _, key := range envKeys {
			fmt.Println(key + ": " + os.Getenv(key))
		}

	}

}

func pingDB() {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoDbServerURI := fmt.Sprintf("%s://%s:%s@%s.%s/?retryWrites=true&w=majority", os.Getenv("MONGODB_SERVER_PREFIX"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("MONGODB_CLUSTER_NAME"), os.Getenv("MONGODB_ENDPOINT_URL"))
	//fmt.Println(mongoDbServerURI)
	opts := options.Client().ApplyURI(mongoDbServerURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Ping to MongoDB Deployment is Successful!")

}
