-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    calorie_intake INT NOT NULL CHECK (calorie_intake >= 0)
);

-- +goose Down
DROP TABLE users;