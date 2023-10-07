// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: posts.sql

package internal

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
	title,
	description,
	category_id,
	path,
	status
) VALUES ($1, $2, $3, $4, $5)
RETURNING id, title, description, category_id, path, status, created_at, updated_at
`

type CreatePostParams struct {
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	CategoryID  uuid.NullUUID  `json:"categoryId"`
	Path        sql.NullString `json:"path"`
	Status      sql.NullString `json:"status"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.Title,
		arg.Description,
		arg.CategoryID,
		arg.Path,
		arg.Status,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CategoryID,
		&i.Path,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCategoryPosts = `-- name: GetCategoryPosts :many
SELECT id, title, description, category_id, path, status, created_at, updated_at FROM posts WHERE category_id=$1
`

func (q *Queries) GetCategoryPosts(ctx context.Context, categoryID uuid.NullUUID) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getCategoryPosts, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CategoryID,
			&i.Path,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPost = `-- name: GetPost :one
SELECT id, title, description, category_id, path, status, created_at, updated_at FROM posts WHERE id=$1 LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CategoryID,
		&i.Path,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :many
SELECT id, title, description, category_id, path, status, created_at, updated_at FROM posts WHERE status=$1
`

func (q *Queries) GetPosts(ctx context.Context, status sql.NullString) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPosts, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CategoryID,
			&i.Path,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTagPosts = `-- name: GetTagPosts :many
SELECT id, title, description, category_id, path, status, created_at, updated_at FROM posts P
WHERE EXISTS (
	SELECT 1 FROM post_tags PT
	WHERE P.id = PT.post_id 
		AND PT.tag_id = $1
)
`

func (q *Queries) GetTagPosts(ctx context.Context, tagID uuid.UUID) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getTagPosts, tagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CategoryID,
			&i.Path,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
