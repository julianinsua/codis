-- name: CreateUser :one
INSERT INTO users (
	username, 
	password,
	email
) VALUES ( $1, $2, $3) 
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id=$1 LIMIT 1;

-- name: GetSessions :one
SELECT * FROM sessions WHERE id=$1 LIMIT 1;
