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

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email=$1 LIMIT 1;

-- name: CreateSession :one
INSERT INTO sessions (
	user_id,
	refresh_token,
	client_agent,
	client_ip,
	is_blocked,
	expires_at,
	created_at
) VALUES ( $1, $2, $3, $4, $5, $6, $7 ) RETURNING *;

-- name: UpdateSession :one
UPDATE sessions
	SET refresh_token =$2, client_agent=$3, client_ip=$4, expires_at=$5
	WHERE id=$1 RETURNING *;
