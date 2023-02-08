package config

import (
	"authentication/pkg/pg"
	"authentication/pkg/token"
	"database/sql"
)

// AppConfig holds the application config
type AppConfig struct {
	DB     *sql.DB
	Models pg.Models
	Paseto *token.Paseto
}
