-- name: CreateIngridient :one
INSERT INTO ingredients (id, created_at, updated_at, calories, name, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetIngridientById :one
SELECT * FROM ingredients WHERE id = $1;