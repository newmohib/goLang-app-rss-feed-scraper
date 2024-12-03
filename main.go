package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	router := chi.NewRouter()

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}
	log.Printf("Server starting on %v", srv.Addr)
	srv.ListenAndServe()
}
