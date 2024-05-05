package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	dbi "github.com/julianinsua/codis/internal/database"
)

type Store struct {
	*dbi.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: dbi.New(db),
	}
}

func (st *Store) execTx(ctx context.Context, fn func(q *dbi.Queries) error) error {
	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Create a new instance of the database code inside using the transaction pointer instead of a database
	q := dbi.New(tx)
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
	Post dbi.Post
	Tags []dbi.Tag
}

func (st Store) CreatePostTx(ctx context.Context, params CreatePostTxParams) (result CreatePostTxResult, err error) {
	err = st.execTx(ctx, func(q *dbi.Queries) error {
		// Create the post
		result.Post, err = q.CreatePost(ctx, dbi.CreatePostParams{
			Title:       params.Title,
			Description: params.Description,
			Status:      sql.NullString{},
			UserID:      params.UserID,
			Path:        params.Path,
		})

		if err != nil {
			return err
		}

		for _, tag := range params.TagNames {
			// TODO: Check if tag exists

			// If it doesn't, create it
			dbTag, err := q.CreateTag(ctx, dbi.CreateTagParams{
				Name:   tag,
				UserID: params.UserID,
			})
			if err != nil {
				return err
			}
			result.Tags = append(result.Tags, dbTag)

			// Connect tags and posts
			_, err = q.CreatePostTag(ctx, dbi.CreatePostTagParams{
				PostID: result.Post.ID,
				TagID:  dbTag.ID,
			})
			if err != nil {
				return err
			}
		}
		return err
	})

	if err != nil {
		return result, fmt.Errorf("unable to execute transaction: %v", err)
	}

	return
}
