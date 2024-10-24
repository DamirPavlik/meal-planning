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
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:            dbUser.ID,
		CreatedAt:     dbUser.CreatedAt,
		UpdatedAt:     dbUser.UpdatedAt,
		Email:         dbUser.Email,
		Password:      dbUser.Password,
		CalorieIntake: dbUser.CalorieIntake,
	}
}
