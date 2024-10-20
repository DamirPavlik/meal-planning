package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/damirpavlik/meal-planning/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email         string `json:"email"`
		Password      string `json:"password"`
		CalorieIntake int32  `json:"calorie_intake"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:            uuid.New(),
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		Email:         params.Email,
		Password:      params.Password,
		CalorieIntake: params.CalorieIntake,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error creating user: %v", err))
		return
	}

	respondWithJSON(w, 201, dbUserToUser(user))
}
