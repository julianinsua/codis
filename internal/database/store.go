package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Store interface {
	Querier
	CreatePostTx(ctx context.Context, params CreatePostTxParams) (result CreatePostTxResult, err error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (st *SQLStore) execTx(ctx context.Context, fn func(q *Queries) error) error {
	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Create a new instance of the database code inside using the transaction pointer instead of a database
	q := New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type CreatePostTxParams struct {
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Status      sql.NullString `json:"status"`
	UserID      uuid.UUID      `json:"userId"`
	Path        string         `json:"path"`
	TagNames    []string       `json:"tagName"`
}

type CreatePostTxResult struct {
	Post Post
	Tags []Tag
}

func (st SQLStore) CreatePostTx(ctx context.Context, params CreatePostTxParams) (result CreatePostTxResult, err error) {
	err = st.execTx(ctx, func(q *Queries) error {
		// Create the post
		result.Post, err = q.CreatePost(ctx, CreatePostParams{
			Title:       params.Title,
			Description: params.Description,
			Status:      params.Status,
			UserID:      params.UserID,
			Path:        params.Path,
		})

		if err != nil {
			return err
		}

		for _, tag := range params.TagNames {
			userTagExistParams := UserTagExistsParams{
				UserID: params.UserID,
				Name:   tag,
			}
			var dbTag Tag
			dbTag, err := q.UserTagExists(ctx, userTagExistParams)
			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}

				dbTag, err = q.CreateTag(ctx, CreateTagParams{
					Name:   tag,
					UserID: params.UserID,
				})
				if err != nil {
					return err
				}
				result.Tags = append(result.Tags, dbTag)
			}

			// Connect tags and posts
			_, err = q.CreatePostTag(ctx, CreatePostTagParams{
				PostID: result.Post.ID,
				TagID:  dbTag.ID,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return result, fmt.Errorf("unable to execute transaction: %v", err)
	}

	return result, err
}
