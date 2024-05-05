package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
	usr := createRandomUser(t)

	createPostParams := CreatePostParams{
		Title: "Test",
		Description: sql.NullString{
			String: "test description",
			Valid:  true,
		},
		Status: sql.NullString{
			String: "Published",
			Valid:  true,
		},
		UserID: usr.ID,
	}

	post, err := testQueries.CreatePost(context.Background(), createPostParams)

	require.NoError(t, err)
	require.Equal(t, post.Title, createPostParams.Title)
	require.Equal(t, post.Description.String, createPostParams.Description.String)
	require.Equal(t, post.Status.String, createPostParams.Status.String)
	require.Equal(t, post.UserID.String(), createPostParams.UserID.String())
}

func TestGetPost(t *testing.T) {
	usr := createRandomUser(t)

	createPostParams := CreatePostParams{
		Title: "Test",
		Description: sql.NullString{
			String: "test description",
			Valid:  true,
		},
		Status: sql.NullString{
			String: "Published",
			Valid:  true,
		},
		UserID: usr.ID,
	}

	post, err := testQueries.CreatePost(context.Background(), createPostParams)
	require.NoError(t, err)

	dbPost, err := testQueries.GetPost(context.Background(), post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, dbPost)
	require.Equal(t, post, dbPost)
}

func TestGetPosts(t *testing.T) {
	usr := createRandomUser(t)

	createPostParams1 := CreatePostParams{
		Title: "Test",
		Description: sql.NullString{
			String: "test description",
			Valid:  true,
		},
		Status: sql.NullString{
			String: "Published",
			Valid:  true,
		},
		UserID: usr.ID,
	}

	post1, err := testQueries.CreatePost(context.Background(), createPostParams1)
	require.NoError(t, err)
	require.NotEmpty(t, post1)

	createPostParams2 := CreatePostParams{
		Title: "Test2",
		Description: sql.NullString{
			String: "test description",
			Valid:  true,
		},
		Status: sql.NullString{
			String: "Published",
			Valid:  true,
		},
		UserID: usr.ID,
	}

	post2, err := testQueries.CreatePost(context.Background(), createPostParams2)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	posts, err := testQueries.GetPosts(context.Background(), sql.NullString{String: "Published", Valid: true})
	require.NoError(t, err)
	require.NotEmpty(t, posts)
	require.GreaterOrEqual(t, len(posts), 2)
}
