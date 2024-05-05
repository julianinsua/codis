-- +goose Up
CREATE TABLE sessions (
    "id" uuid PRIMARY KEY,
    "user_id" uuid NOT NULL UNIQUE,
    "refresh_token" varchar NOT NULL,
    "client_agent" varchar NOT NULL,
    "client_ip" varchar NOT NULL,
    "is_blocked" boolean NOT NULL DEFAULT FALSE,
    "expires_at" timestamp NOT NULL,
    "created_at" timestamp NOT NULL
);

ALTER TABLE sessions ADD FOREIGN KEY ("user_id") REFERENCES users ("id");

-- +goose Down
DROP TABLE IF EXISTS sessions;
