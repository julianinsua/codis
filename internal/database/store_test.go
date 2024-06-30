package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/julianinsua/codis/util"
	"github.com/stretchr/testify/require"
)

type comparison struct {
	Expected CreatePostTxParams
	Result   CreatePostTxResult
}

func TestCreatePostTx(t *testing.T) {
	store := NewStore(testDB)
	user := CreateRandomUser(t)

	errs := make(chan error)
	txResults := make(chan comparison)

	n := 5
	for i := 0; i < n; i++ {
		go func() {
			params := CreatePostTxParams{
				Title: util.RandomString(6),
				Description: sql.NullString{
					String: util.RandomString(12),
					Valid:  true,
				},
				Status: sql.NullString{
					String: "Published",
					Valid:  true,
				},
				UserID:   user.ID,
				Path:     "path/to/file",
				TagNames: []string{util.RandomString(5), util.RandomString(7), util.RandomString(9)},
			}

			result, err := store.CreatePostTx(context.Background(), params)
			errs <- err
			txResults <- comparison{Expected: params, Result: result}
		}()
	}

	// Listen for results and errors
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		res := <-txResults
		require.NotEmpty(t, res)
		require.NotEmpty(t, res.Result)
		require.NotEmpty(t, res.Expected)

		require.Equal(t, res.Expected.Title, res.Result.Post.Title)
		require.Equal(t, res.Expected.Description.String, res.Result.Post.Description.String)
		require.Equal(t, res.Expected.Status.String, res.Result.Post.Status.String)
		require.Equal(t, res.Expected.UserID, res.Result.Post.UserID)
		require.Equal(t, res.Expected.Path, res.Result.Post.Path)

		for _, tag := range res.Result.Tags {
			require.Contains(t, res.Expected.TagNames, tag.Name)
		}
	}
}
