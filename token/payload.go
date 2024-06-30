package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
)

// The basic struct that holds the tokens related to the claims in a token
type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

// Holds the payload that is encoded inside the PASETO token
type PASETOPayload struct {
	Payload
	ExpiresAt time.Time `json:"expiresAt"`
	IssuedAt  time.Time `json:"IssuedAt"`
}

// Creates a new PASETO token payload using username and duration.
func NewPASETOPayload(username string, duration time.Duration) (*PASETOPayload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &PASETOPayload{Payload{tokenId, username}, time.Now().Add(duration), time.Now()}
	return payload, nil
}
