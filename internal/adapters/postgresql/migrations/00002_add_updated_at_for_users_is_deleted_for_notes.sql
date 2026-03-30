-- +goose Up
ALTER TABLE users
ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT now();

ALTER TABLE notes
ADD COLUMN is_deleted BOOLEAN NOT NULL DEFAULT false;

-- +goose Down
ALTER TABLE
DROP COLUMN is_deleted;

ALTER TABLE users
DROP COLUMN updated_at;

