package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/newmohib/goLang-app-rss-feed-scraper/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type createUserRequest struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := createUserRequest{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:   uuid.New(),
		Name: params.Name,
		// CreatedAt: time.UTC(),
		// UpdatedAt: time.UTC(),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to create user: %v", err))
		return
	}

	responsdWithJSON(w, http.StatusOK, databaseUserToUser(user))

}
