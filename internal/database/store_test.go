package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/julianinsua/codis/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePostTx(t *testing.T) {
	store := NewStore(testDB)
	user := createRandomUser(t)

	errs := make(chan error)
	txResults := make(chan CreatePostTxResult)

	n := 5
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreatePostTx(context.Background(), CreatePostTxParams{
				Title: util.RandomString(6),
				Description: sql.NullString{
					String: util.RandomString(12),
					Valid:  true,
				},
				Status: sql.NullString{
					String: util.RandomString(5),
					Valid:  true,
				},
				UserID:   user.ID,
				Path:     "path/to/file", // TODO: Remove the path, it no longer makes sense
				TagNames: []string{util.RandomString(5), util.RandomString(7), util.RandomString(9)},
			})
			errs <- err
			txResults <- result
		}()
	}

	// Listen for results and errors
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		res := <-txResults
		require.NotEmpty(t, res)

		// TODO: further checks here
	}
}
