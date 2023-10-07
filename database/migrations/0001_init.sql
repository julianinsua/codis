-- +goose Up
CREATE TABLE posts (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "title" varchar NOT NULL,
  "description" varchar,
  "category_id" uuid,
  "path" varchar UNIQUE,
  "status" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE categories (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE post_tags (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "post_id" uuid NOT NULL,
  "tag_id" uuid NOT NULL
);

CREATE TABLE tags (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar
);

CREATE INDEX ON posts ("title");

CREATE INDEX ON posts ("description");

CREATE INDEX ON posts ("category_id");

CREATE INDEX ON posts ("created_at");

CREATE INDEX ON posts ("updated_at");

ALTER TABLE posts ADD FOREIGN KEY ("category_id") REFERENCES categories ("id");

ALTER TABLE post_tags ADD FOREIGN KEY ("post_id") REFERENCES posts ("id");

ALTER TABLE post_tags ADD FOREIGN KEY ("tag_id") REFERENCES tags ("id");

-- +goose Down
DROP TABLE posts;
DROP TABLE categories;
DROP TABLE post_tags;
DROP TABLE tags;
