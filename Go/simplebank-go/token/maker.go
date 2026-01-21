package token

import (
	"time"
)

// Makaer is an interface for managing tokens
type Maker interface {
	// create and sign a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// check if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
