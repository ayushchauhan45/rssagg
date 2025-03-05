package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ayushchauhan_45/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	feedFollow, er := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if er != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating feedfollow: %v", er))
		return
	}

	respondWithJson(w, 201, databaseFeedFollowtoFeedFollow(feedFollow))
}

func (apiCfg *apiconfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollow, er := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if er != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follow: %v", er))
		return
	}

	respondWithJson(w, 201, databaseFeedFollowstoFeedFollows(feedFollow))
}
func (apiCfg *apiconfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowId, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parse feed follow id: %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	respondWithJson(w, 200, struct{}{})
}
