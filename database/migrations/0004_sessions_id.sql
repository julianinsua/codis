-- +goose Up
ALTER TABLE sessions ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- +goose Down
ALTER TABLE sessions ALTER COLUMN id DROP DEFAULT;
