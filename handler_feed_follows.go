package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/newmohib/goLang-app-rss-feed-scraper/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type createUserRequest struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := createUserRequest{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to create Feed: %v", err))
		return
	}

	responsdWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))

}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not get Feed follows: %v", err))
		return
	}

	responsdWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))

}


func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing feedFollowID: %v", err))
		return
	}

	err2 := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		FeedID:     feedFollowID,
		UserID: user.ID,
	})

	if err2 != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not Delete Feed follow: %v", err2))
		return
	}

	responsdWithJSON(w, http.StatusOK, struct{}{})

}