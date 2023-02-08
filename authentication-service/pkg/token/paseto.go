package token

import (
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"

	"github.com/o1egl/paseto"
)

// Paseto is a PASETO token type
type Paseto struct {
	paseto      *paseto.V2
	symetricKey []byte
}

// NewPaseto creates a new Paseto token
func NewPaseto(symetricKey string) (*Paseto, error) {
	if len(symetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &Paseto{
		paseto:      paseto.NewV2(),
		symetricKey: []byte(symetricKey),
	}
	return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (p *Paseto) CreateToken(email, nickname string, role int) (string, error) {
	payload, err := NewPayload(email, nickname, role)
	if err != nil {
		return "", err
	}

	return p.paseto.Encrypt(p.symetricKey, payload, nil)
}

// DecryptToken decrypts the token into TokenData
func (p *Paseto) DecryptToken(token string) (*TokenData, error) {
	td := &TokenData{}

	err := p.paseto.Decrypt(token, p.symetricKey, td, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	return td, nil
}
