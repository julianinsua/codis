-- name: CreateTag :one
INSERT INTO tags (
	 name, user_id
) VALUES ( $1, $2 ) RETURNING *;

-- name: GetUserTags :many
SELECT * FROM tags WHERE user_id=$1;

-- name: GetTagById :one
SELECT * FROM tags WHERE id=$1;
