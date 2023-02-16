package config

import (
	"database/sql"
	"employee/pkg/pg"
)

// AppConfig holds the application config
type AppConfig struct {
	DB     *sql.DB
	Models pg.Models
}
