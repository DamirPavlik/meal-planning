-- +goose Up
ALTER TABLE ingredients ADD COLUMN user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE ingredients DROP COLUMN user_id;