package handlers

import (
	"frontend/pkg/config"
)

// CookieName is the name of the secure cookie which holds the paseto token
const CookieName = "authToken"

var (
	httpContext httpContextStruct // httpContext is the context type used by the middleware
	Repo        *Repository       // Repo is the repository used by the handlers
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
