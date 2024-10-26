-- +goose Up
ALTER TABLE users ADD bearers_token VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex')   
)

-- +goose Down
DROP TABLE users;