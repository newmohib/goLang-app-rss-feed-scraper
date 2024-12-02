package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	myEnv, err := godotenv.Read()
	portString := myEnv["PROT"]
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Hello World")
	fmt.Println(portString)
}
