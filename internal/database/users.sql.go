// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
	user_id,
	refresh_token,
	client_agent,
	client_ip,
	is_blocked,
	expires_at,
	created_at
) VALUES ( $1, $2, $3, $4, $5, $6, $7 ) RETURNING id, user_id, refresh_token, client_agent, client_ip, is_blocked, expires_at, created_at
`

type CreateSessionParams struct {
	UserID       uuid.UUID `json:"userId"`
	RefreshToken string    `json:"refreshToken"`
	ClientAgent  string    `json:"clientAgent"`
	ClientIp     string    `json:"clientIp"`
	IsBlocked    bool      `json:"isBlocked"`
	ExpiresAt    time.Time `json:"expiresAt"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.UserID,
		arg.RefreshToken,
		arg.ClientAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiresAt,
		arg.CreatedAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.ClientAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
	username, 
	password,
	email
) VALUES ( $1, $2, $3) 
RETURNING id, username, password, email, created_at, updated_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSessions = `-- name: GetSessions :one
SELECT id, user_id, refresh_token, client_agent, client_ip, is_blocked, expires_at, created_at FROM sessions WHERE id=$1 LIMIT 1
`

func (q *Queries) GetSessions(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessions, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.ClientAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, password, email, created_at, updated_at FROM users WHERE email=$1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, password, email, created_at, updated_at FROM users WHERE id=$1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSession = `-- name: UpdateSession :one
UPDATE sessions
	SET refresh_token =$2, client_agent=$3, client_ip=$4, expires_at=$5
	WHERE id=$1 RETURNING id, user_id, refresh_token, client_agent, client_ip, is_blocked, expires_at, created_at
`

type UpdateSessionParams struct {
	ID           uuid.UUID `json:"id"`
	RefreshToken string    `json:"refreshToken"`
	ClientAgent  string    `json:"clientAgent"`
	ClientIp     string    `json:"clientIp"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, updateSession,
		arg.ID,
		arg.RefreshToken,
		arg.ClientAgent,
		arg.ClientIp,
		arg.ExpiresAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.ClientAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
