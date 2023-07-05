package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/SEANYB4/blog_aggregator/internal/database"
	_ "github.com/SEANYB4/blog_aggregator/internal/auth"
	"github.com/google/uuid"
	"time"

)



func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

		type parameters struct {
			Name string `json:"name"`
		}

		decoder := json.NewDecoder(r.Body)

		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
			return
		}


		user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name: params.Name,

		})

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %s", err))
			return
		}


		respondWithJSON(w, 201, databaseUserToUser(user))
}



func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {


		respondWithJSON(w, 200, databaseUserToUser(user))
	
}