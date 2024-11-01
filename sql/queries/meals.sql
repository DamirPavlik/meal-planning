-- name: CreateMeal :one
INSERT INTO meals (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: AddIngredientsToMeal :exec
INSERT INTO meal_ingredients (meal_id, ingredient_id)
SELECT $1, unnest($2::UUID[]);