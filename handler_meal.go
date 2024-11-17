package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/damirpavlik/meal-planning/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateMeal(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if params.Name == "" {
		respondWithError(w, 400, "no name found")
		return
	}
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed parsing json: %v", err))
		return
	}

	meal, err := apiCfg.DB.CreateMeal(r.Context(), database.CreateMealParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to create a meal: %v", err))
		return
	}

	respondWithJSON(w, 200, dbMealToMeal(meal))

}
