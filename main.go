package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/christus02/worklog/db"
	"github.com/joho/godotenv"
)

func init() {
	// Load the ENV variables of this project
	loadEnv("worklog.env", false)

}

func main() {

	dbConnection := db.New(
		os.Getenv("MONGODB_SCHEME"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("MONGODB_CLUSTER_NAME"),
		os.Getenv("MONGODB_SERVER_ENDPOINT"),
		os.Getenv("WORKLOG_DB_NAME"),
	)
	dbConnection.Connect()
	dbConnection.PingDB()

	newTask := db.Task{
		Tag:         "test",
		Summary:     "Test Summary 2",
		Comments:    []string{"First Comment"},
		CreatedAt:   time.Now(),
		LastUpdated: time.Now(),
	}
	dbConnection.InsertNewTask(newTask)

	dbConnection.Disconnect()

}

func loadEnv(fileName string, verbose bool) {
	fmt.Println("Loading ENV for Application")
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
