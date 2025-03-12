package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ayushchauhan_45/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	users, er := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if er != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", er))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(users))
}

func (apiCfg *apiconfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiconfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostForUser(r.Context(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating post: %v", err))
		return
	}
	respondWithJson(w, 200, databasePoststoPosts(posts))
}
