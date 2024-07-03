package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/julianinsua/codis/util"
)

func RandomUser() (User, string, error) {
	pass := util.RandomPassword()
	hash, err := util.HashPassword(pass)
	if err != nil {
		return User{}, "", err
	}
	return User{
		ID:        uuid.New(),
		Username:  util.RandomUsername(),
		Password:  hash,
		Email:     util.RandomEmail(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}, pass, nil
}
