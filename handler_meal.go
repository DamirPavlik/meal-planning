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

	mealNames, err := apiCfg.DB.GetAllMealsForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed getting meals: %v", err))
		return
	}

	for _, meal := range mealNames {
		if meal == params.Name {
			respondWithError(w, 400, "user already has a meal with that name")
			return
		}
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

func (apiCfg *apiConfig) handlerGetMealByID(w http.ResponseWriter, r *http.Request, user database.User) {
	mealIDstr := chi.URLParam(r, "mealId")
	if mealIDstr == "" {
		respondWithError(w, 400, "meal id is empty")
		return
	}

	mealID, err := uuid.Parse(mealIDstr)
	if err != nil {
		respondWithError(w, 400, "could not parse meal id")
		return
	}

	meal, err := apiCfg.DB.GetMealById(r.Context(), mealID)
	if err != nil {
		respondWithError(w, 400, "could not get a meal")
		return
	}

	respondWithJSON(w, 200, meal)
}
