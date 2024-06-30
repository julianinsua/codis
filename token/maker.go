package token

import "time"

// Interface to manage tokens
type Maker interface {
	// Creates a new tokenfor a specific user and for a given duration
	CreateToken(username string, duration time.Duration) (string, *PASETOPayload, error)
	// Verifies a token string
	VerifyToken(token string) (*PASETOPayload, error)
}
