package util

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/julianinsua/codis/internal/database"
)

const (
	alpha = "abcdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Returns a random integer in the range provided by min - max. It returns an error if the range length is 0 or negative (min >= max).
func RandomInt(min, max int64) (int64, error) {
	if min >= max {
		return 0, errors.New("invalid range")
	}
	return min + rand.Int63n(max-min), nil
}

// Returns a random string composed of uppercase and lowercase characters from the english language
func RandomString(length int) string {
	var sb strings.Builder
	k := len(alpha)

	for i := 1; i < length; i++ {
		c := alpha[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Returns a random name composed of 6 letters for the Owner of an account
func RandomUsername() string {
	return RandomString(6)
}

// Returns a random email account using random characters for both the user and the domain. The suffix is ".com" by default.
func RandomEmail() string {
	return fmt.Sprintf("%s@%s.com", RandomString(6), RandomString(6))
}

func RandomPassword() string {
	return RandomString(8)
}

func RandomUser() (database.User, string, error) {
	pass := RandomPassword()
	hash, err := HashPassword(pass)
	if err != nil {
		return database.User{}, "", err
	}
	return database.User{
		ID:        uuid.New(),
		Username:  RandomUsername(),
		Password:  hash,
		Email:     RandomEmail(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}, pass, nil
}
