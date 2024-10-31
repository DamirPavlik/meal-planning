package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/damirpavlik/meal-planning/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateIngridient(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Calories int    `json:"calories"`
		Name     string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("err parsing json: %v", err))
	}

	ingridient, err := apiCfg.DB.CreateIngridient(r.Context(), database.CreateIngridientParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Calories:  int32(params.Calories),
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error creating ingridient: %v", err))
		return
	}

	respondWithJSON(w, 200, dbIngridientToIngridient(ingridient))
}
