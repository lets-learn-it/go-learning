package token

import "time"

// it will make switching token version/type easy
type Maker interface {
	CreateToken(username string, duration time.Duration, scopes []string) (string, error)

	VerifyToken(token string, scopes []string) (*Payload, error)
}
