package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	// "github.com/newmohib/goLang-app-rss-feed-scraper/internal/auth"
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

// GET User Handler by apikey

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// //check authentication header
	// apiKey, err := auth.GetAPIKey(r.Header)
	// if err != nil {
	// 	respondWithError(w, http.StatusUnauthorized, err.Error())
	// 	return
	// }
	// //get user
	// user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user: %v", err))
	// 	return
	// }
	responsdWithJSON(w, http.StatusOK, databaseUserToUser(user))

}
func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  500,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not get posts: %v", err))
		return
	}
	responsdWithJSON(w, http.StatusOK, databasePostsToPosts(posts))

}
