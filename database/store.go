package db

import (
	"context"
	"database/sql"
	"fmt"

	dbi "github.com/julianinsua/codis/database/internal"
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
