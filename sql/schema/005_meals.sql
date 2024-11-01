-- +goose Up
CREATE TABLE meals (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name TEXT NOT NULL
);

CREATE TABLE meal_ingredients (
    meal_id UUID REFERENCES meals(id) ON DELETE CASCADE,
    ingredient_id UUID REFERENCES ingredients(id) ON DELETE CASCADE,
    PRIMARY KEY (meal_id, ingredient_id)
);

CREATE VIEW meal_calories AS
SELECT
    m.id AS meal_id,
    m.name,
    m.created_at,
    m.updated_at,
    COALESCE(SUM(i.calories), 0) AS calories
FROM
    meals m
LEFT JOIN
    meal_ingredients mi ON m.id = mi.meal_id
LEFT JOIN
    ingredients i ON mi.ingredient_id = i.id
GROUP BY
    m.id;

-- +goose Down
DROP VIEW meal_calories;
DROP TABLE meal_ingredients;
DROP TABLE meals;
