// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ingridients.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createIngridient = `-- name: CreateIngridient :one
INSERT INTO ingredients (id, created_at, updated_at, calories, name, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, calories, name, user_id
`

type CreateIngridientParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Calories  int32
	Name      string
	UserID    uuid.UUID
}

func (q *Queries) CreateIngridient(ctx context.Context, arg CreateIngridientParams) (Ingredient, error) {
	row := q.db.QueryRowContext(ctx, createIngridient,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Calories,
		arg.Name,
		arg.UserID,
	)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Calories,
		&i.Name,
		&i.UserID,
	)
	return i, err
}

const getIngridientById = `-- name: GetIngridientById :one
SELECT id, created_at, updated_at, calories, name, user_id FROM ingredients WHERE id = $1
`

func (q *Queries) GetIngridientById(ctx context.Context, id uuid.UUID) (Ingredient, error) {
	row := q.db.QueryRowContext(ctx, getIngridientById, id)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Calories,
		&i.Name,
		&i.UserID,
	)
	return i, err
}
