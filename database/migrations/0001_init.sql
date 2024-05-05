-- +goose Up
CREATE TABLE posts (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "title" varchar NOT NULL,
  "description" varchar,
  "status" varchar,
  "user_id" uuid NOT NULL,
  "path" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE post_tags (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "post_id" uuid NOT NULL,
  "tag_id" uuid NOT NULL
);

CREATE TABLE tags (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE users (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);


CREATE INDEX ON posts ("title");

CREATE INDEX ON posts ("description");

CREATE INDEX ON posts ("created_at");

CREATE INDEX ON posts ("updated_at");

CREATE INDEX ON posts ("user_id");

ALTER TABLE posts ADD FOREIGN KEY ("user_id") references users ("id");

ALTER TABLE post_tags ADD FOREIGN KEY ("post_id") REFERENCES posts ("id");

ALTER TABLE post_tags ADD FOREIGN KEY ("tag_id") REFERENCES tags ("id");

ALTER TABLE tags ADD FOREIGN KEY ("user_id") REFERENCES users ("id");

-- +goose Down
DROP TABLE post_tags;
DROP TABLE users;
DROP TABLE posts;
DROP TABLE tags;
