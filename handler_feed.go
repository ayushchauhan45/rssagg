package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ayushchauhan_45/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	feed, er:= apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if  er != nil{
		respondWithError(w, 400, fmt.Sprintf("Error creating feed: %v", er))
		return
	}

	respondWithJson(w, 201, databaseFeedtoFeed(feed))
}

func (apiCfg *apiconfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {
	feed, er := apiCfg.DB.GetFeeds(r.Context())
	if  er != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed: %v", er))
		return
	}

	respondWithJson(w, 201, databaseFeedstoFeeds(feed))
}

