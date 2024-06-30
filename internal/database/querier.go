// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreatePostTag(ctx context.Context, arg CreatePostTagParams) (PostTag, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetPost(ctx context.Context, id uuid.UUID) (Post, error)
	GetPosts(ctx context.Context, status sql.NullString) ([]Post, error)
	// need to add ordering
	GetPostsWithTags(ctx context.Context, userID uuid.UUID) ([]PostsView, error)
	GetSessions(ctx context.Context, id uuid.UUID) (Session, error)
	GetTagById(ctx context.Context, id uuid.UUID) (Tag, error)
	GetTagPosts(ctx context.Context, tagID uuid.UUID) ([]Post, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserTags(ctx context.Context, userID uuid.UUID) ([]Tag, error)
	UpdateSession(ctx context.Context, arg UpdateSessionParams) (Session, error)
	UserTagExists(ctx context.Context, arg UserTagExistsParams) (Tag, error)
}

var _ Querier = (*Queries)(nil)
