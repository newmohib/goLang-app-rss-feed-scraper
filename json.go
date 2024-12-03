package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string)  {
	if code > 499 {
		log.Println("Responding with 5xx error to the client error:", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	responsdWithJSON(w, code, errorResponse{Error: msg})
}

func responsdWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}