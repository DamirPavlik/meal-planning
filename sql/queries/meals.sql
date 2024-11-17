-- name: CreateMeal :one
INSERT INTO meals (id, created_at, updated_at, name, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *; 

-- name: AddIngredientsToMeal :exec
INSERT INTO meal_ingredients (meal_id, ingredient_id)
SELECT $1, unnest($2::UUID[]);

-- name: GetAllMealsForUser :many
SELECT name FROM meals WHERE user_id = $1;