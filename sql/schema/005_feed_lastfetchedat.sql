-- +goose Up
ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
Alter TABLE DROP COLUMN last_fetched_at;
