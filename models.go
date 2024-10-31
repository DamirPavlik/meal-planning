package main

import (
	"time"

	"github.com/damirpavlik/meal-planning/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	CalorieIntake int32     `json:"calorie_intake"`
	BearersToken  string    `json:"bearers_token"`
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:            dbUser.ID,
		CreatedAt:     dbUser.CreatedAt,
		UpdatedAt:     dbUser.UpdatedAt,
		Email:         dbUser.Email,
		Password:      dbUser.Password,
		CalorieIntake: dbUser.CalorieIntake,
		BearersToken:  dbUser.BearersToken,
	}
}

type Ingridient struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at`
	Name      string    `json:"name"`
	Calories  int       `json:"calories"`
	UserID    uuid.UUID `json:"user_id"`
}

func dbIngridientToIngridient(dbIngridient database.Ingredient) Ingridient {
	return Ingridient{
		ID:        dbIngridient.ID,
		CreatedAt: dbIngridient.CreatedAt,
		UpdatedAt: dbIngridient.UpdatedAt,
		Name:      dbIngridient.Name,
		Calories:  int(dbIngridient.Calories),
		UserID:    dbIngridient.UserID,
	}
}
