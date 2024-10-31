package main

import (
	"fmt"
	"net/http"

	"github.com/damirpavlik/meal-planning/internal/auth"
	"github.com/damirpavlik/meal-planning/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := auth.GetBearersToken(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("auth err: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByBearers(r.Context(), key)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("could not get user: %v", err))
			return
		}
		handler(w, r, user)

	}
}
