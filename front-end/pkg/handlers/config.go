package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

// VerifyCookiePayload is the payload send to authentication-service
// to verify the cookie & token
type VerifyCookiePayload struct {
	Cookie http.Cookie `json:"cookie"`
	// Cookie string `json:"cookie"`
}

// ResponseCookieVerify is the response sent from authentication-service
type ResponseCookieVerify struct {
	Status       int       `json:"status"`
	Message      string    `json:"message"`
	TokenPayload TokenData `json:"tokenpayload"`
}

// TokenData contains the payload data of the token
type TokenData struct {
	ID        uuid.UUID `json:"uuid"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Role      int       `json:"role"`
	IssuedAt  time.Time `json:"issue_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// httpContextStruct contains the type to be sent
// to the context by the middleware
type httpContextStruct struct {
	Auth bool
	User TokenData
}
