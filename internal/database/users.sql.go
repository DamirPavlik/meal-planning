// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, password, calorie_intake, bearers_token)
VALUES ($1, $2, $3, $4, $5, $6,
    encode(sha256(random()::text::bytea), 'hex')   
)
RETURNING id, created_at, updated_at, email, password, calorie_intake, bearers_token
`

type CreateUserParams struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Email         string
	Password      string
	CalorieIntake int32
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Email,
		arg.Password,
		arg.CalorieIntake,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Password,
		&i.CalorieIntake,
		&i.BearersToken,
	)
	return i, err
}

const getUserByBearers = `-- name: GetUserByBearers :one
SELECT id, created_at, updated_at, email, password, calorie_intake, bearers_token FROM users WHERE bearers_token = $1
`

func (q *Queries) GetUserByBearers(ctx context.Context, bearersToken string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByBearers, bearersToken)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Password,
		&i.CalorieIntake,
		&i.BearersToken,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, created_at, updated_at, email, password, calorie_intake, bearers_token FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Password,
		&i.CalorieIntake,
		&i.BearersToken,
	)
	return i, err
}
