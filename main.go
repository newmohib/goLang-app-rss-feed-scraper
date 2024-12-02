package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// get environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	  }
	  
	portString := os.Getenv("PORT")
	fmt.Println("Hello World", portString)
}
