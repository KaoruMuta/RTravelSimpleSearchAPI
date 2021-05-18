package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)
	
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to loading .env file")
	}

	applicationId := os.Getenv("APPLICATION_ID")
	fmt.Println(applicationId)
}