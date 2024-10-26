-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, password, calorie_intake, bearers_token)
VALUES ($1, $2, $3, $4, $5, 
    encode(sha256(random()::text::bytea), 'hex')   
)
RETURNING *;