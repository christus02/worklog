package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("worklog.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPasswd := os.Getenv("DB_PASSWORD")

	log.Println("DB Username: " + dbUser + " | DB Password: " + dbPasswd)

}
