package main

import (
	"log"
	"mailer/pkg/config"
	"mailer/pkg/handlers"
	"mailer/pkg/web"
)

const webPort = "8888"

func main() {
	// app holds our application configuration
	var app config.AppConfig

	// Customize log display
	setLog()

	// set the config for the handlers package
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Start web server
	if err := web.Run(multiplexer(), webPort); err != nil {
		log.Fatalf(err.Error())
	}
}
