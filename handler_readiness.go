package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responsdWithJSON(w, http.StatusOK, struct{}{})
	
}