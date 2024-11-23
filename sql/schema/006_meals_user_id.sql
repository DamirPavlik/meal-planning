-- +goose Up
ALTER TABLE meals ADD COLUMN user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE meals DROP COLUMN user_id;