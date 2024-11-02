package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/damirpavlik/meal-planning/internal/database"
	"github.com/go-chi/chi"
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

func (apiCfg *apiConfig) handlerGetUserByID(w http.ResponseWriter, r *http.Request, user database.User) {
	userIDStr := chi.URLParam(r, "userID")
	if userIDStr == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user ID format", http.StatusBadRequest)
		return
	}

	retrievedUser, err := apiCfg.DB.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, dbUserToUser(retrievedUser))
}
