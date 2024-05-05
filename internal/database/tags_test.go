package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/julianinsua/codis/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTag(t *testing.T) Tag {
	usr := createRandomUser(t)

	params := CreateTagParams{
		Name:   util.RandomString(6),
		UserID: usr.ID,
	}

	tag, err := testQueries.CreateTag(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, tag)
	require.Equal(t, params.Name, tag.Name)
	require.Equal(t, params.UserID, usr.ID)

	return tag
}

func TestCreateTag(t *testing.T) {
	CreateRandomTag(t)
}

func TestGetUserTags(t *testing.T) {
	usr := createRandomUser(t)

	params1 := CreateTagParams{
		Name:   util.RandomString(6),
		UserID: usr.ID,
	}
	tag1, err := testQueries.CreateTag(context.Background(), params1)
	require.NoError(t, err)
	require.NotEmpty(t, tag1)

	params2 := CreateTagParams{
		Name:   util.RandomString(6),
		UserID: usr.ID,
	}
	tag2, err := testQueries.CreateTag(context.Background(), params2)
	require.NoError(t, err)
	require.NotEmpty(t, tag2)

	tags, err := testQueries.GetUserTags(context.Background(), usr.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tags)

	require.Equal(t, tags[0], tag1)
	require.Equal(t, tags[1], tag2)
}

func TestGetTagById(t *testing.T) {
	tag := CreateRandomTag(t)

	dbTag, err := testQueries.GetTagById(context.Background(), tag.ID)
	require.NoError(t, err)
	require.NotEmpty(t, dbTag)
	require.Equal(t, dbTag, tag)
}

func TestGetPostsWithTags(t *testing.T) {
	usr := createRandomUser(t)
	tagParams := CreateTagParams{
		Name:   util.RandomString(6),
		UserID: usr.ID,
	}
	tag, err := testQueries.CreateTag(context.Background(), tagParams)
	require.NoError(t, err)
	require.NotEmpty(t, tag)

	postParams := CreatePostParams{
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
	post, err := testQueries.CreatePost(context.Background(), postParams)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	postTagParams := CreatePostTagParams{
		PostID: post.ID,
		TagID:  tag.ID,
	}
	posts, err := testQueries.CreatePostTag(context.Background(), postTagParams)
	require.NoError(t, err)
	require.NotEmpty(t, posts)
	require.NotEmpty(t, posts.ID)
	require.Equal(t, posts.TagID, tag.ID)
	require.Equal(t, posts.PostID, post.ID)

	postsWithTags, err := testQueries.GetPostsWithTags(context.Background(), usr.ID)
	require.NoError(t, err)
	require.NotEmpty(t, postsWithTags)

	var dbTags []Tag
	dbTagJson := postsWithTags[0].PostTags
	err = json.Unmarshal(dbTagJson, &dbTags)
	require.NoError(t, err)

	require.Equal(t, tag.ID, dbTags[0].ID)
	require.Equal(t, tag.Name, dbTags[0].Name)
	require.Equal(t, tag.UserID, dbTags[0].UserID)
}
