package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// TokenData contains the payload data of the token
type TokenData struct {
	ID        uuid.UUID `json:"uuid"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Role      int       `json:"role"`
	IssuedAt  time.Time `json:"issue_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token for a specific username (email, nickname, role)
func NewPayload(email, nickname string, role int) (*TokenData, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	p := &TokenData{
		ID:        tokenID,
		Email:     email,
		Nickname:  nickname,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(Duration),
	}

	return p, nil
}

// Valid checks if the token payload is expired
func (td *TokenData) Valid() error {
	if time.Now().After(td.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}

//  Valid checks if the token payload is expired
// returns duration left
// func (td *TokenData) Valid() (time.Duration, error) {
// 	if time.Now().After(td.ExpiredAt) {
// 		return 0, ErrExpiredToken
// 	}
// 	// should use time.Until instead of t.Sub(time.Now())
// 	duration := td.ExpiredAt.Sub(time.Now())

//		return duration, nil
//	}
