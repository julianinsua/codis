package database

import (
	"context"
	"testing"

	"github.com/julianinsua/codis/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	hash, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	params := CreateUserParams{
		Username: util.RandomUsername(),
		Password: hash,
		Email:    util.RandomEmail(),
	}

	usr, err := testQueries.CreateUser(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, usr)
	require.Equal(t, params.Username, usr.Username)
	require.Equal(t, params.Email, usr.Email)
	require.Equal(t, params.Password, usr.Password)
	require.NotZero(t, usr.CreatedAt)
	require.NotZero(t, usr.UpdatedAt)

	return usr
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetuserByID(t *testing.T) {
	usr := CreateRandomUser(t)

	dbUsr, err := testQueries.GetUserByID(context.Background(), usr.ID)

	require.NoError(t, err)
	require.Equal(t, usr.ID, dbUsr.ID)
	require.Equal(t, usr.Username, dbUsr.Username)
	require.Equal(t, usr.Password, dbUsr.Password)
	require.Equal(t, usr.Email, dbUsr.Email)
	require.Equal(t, usr.CreatedAt, dbUsr.CreatedAt)
	require.Equal(t, usr.UpdatedAt, dbUsr.UpdatedAt)
}
