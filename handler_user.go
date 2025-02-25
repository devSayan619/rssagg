package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devSayan619/rssagg/internal/database"
	"github.com/google/uuid"
)

func(apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err :=	decoder.Decode(&params)
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user,err :=	apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name: params.Name,
	})

	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Couldn't create user:%v",err))
		return
	}
	respondWithJSON(w, 201, databaseUserToUser(user))

}

func(apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request,user database.User) {
	respondWithJSON(w,200,databaseUserToUser(user))
}

func(apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request,user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})

	if err!=nil {
		respondWithError(w,400,fmt.Sprintf("couldn't get posts: %v",err))
		return
	}

	respondWithJSON(w,200,databasePostsToPosts(posts))
}