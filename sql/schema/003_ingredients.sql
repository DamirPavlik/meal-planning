-- +goose Up

CREATE TABLE ingredients (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    calories INT NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE ingredients;