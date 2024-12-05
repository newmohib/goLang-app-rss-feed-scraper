package main

import (
	"fmt"
	"net/http"

	"github.com/newmohib/goLang-app-rss-feed-scraper/internal/auth"
	"github.com/newmohib/goLang-app-rss-feed-scraper/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check authentication header
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		//get user
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
