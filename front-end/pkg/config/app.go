package config

import (
	"html/template"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool // Used for development purpose (When set to false can refresh the html page without recompiling app)
	TemplateCache map[string]*template.Template
}
