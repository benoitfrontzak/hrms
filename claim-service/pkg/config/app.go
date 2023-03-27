package config

import (
	"claim/pkg/pg"
	"database/sql"
)

// AppConfig holds the application config
type AppConfig struct {
	DB     *sql.DB
	Models pg.Models
}
