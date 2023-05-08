package handlers

import (
	"mime/multipart"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// VerifyCookiePayload is the payload send to authentication-service
// to verify the cookie & token
type VerifyCookiePayload struct {
	Cookie http.Cookie `json:"cookie"`
}

// ResponseCookieVerify is the response sent from authentication-service
type ResponseCookieVerify struct {
	Status       int       `json:"status"`
	Message      string    `json:"message"`
	TokenPayload TokenData `json:"tokenpayload"`
}

// TokenData contains the payload data of the token
type TokenData struct {
	ID         uuid.UUID `json:"uuid"`
	Email      string    `json:"email"`
	Nickname   string    `json:"nickname"`
	Role       int       `json:"role"`
	EmployeeID int       `json:"employeeID"`
	IssuedAt   time.Time `json:"issue_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}

// httpContextStruct contains the type to be sent
// to the context by the middleware
type httpContextStruct struct {
	Auth bool
	User TokenData
}

// TemplateData is the struct given to any HTML template
type TemplateData struct {
	User TokenData
	Data any
}

// jsonResponse is a struct which holds the response to be sent back
type jsonResponse struct {
	Error     bool   `json:"error"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
}

// attachments is the struct which holds the payload from upload files form
type attachments struct {
	Files         []*multipart.FileHeader
	Filename      string
	ApplicationID string
	Email         string
	ID            string
}

// uploadedFiles is the struct which holds all uploaded files of one employee
type uploadedFiles struct {
	Files   map[string][]string
	Archive map[string]map[string][]string
}
