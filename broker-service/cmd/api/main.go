package main

import (
	"broker/pkg/config"
	"broker/pkg/handlers"
	"broker/pkg/rabbitmq"
	"broker/pkg/web"
	"log"
	"os"
)

const webPort = "8088"

func main() {
	// app holds our application configuration
	var app config.AppConfig

	// connect to rabbitmq server
	rabbitConn, err := rabbitmq.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	// set the config for the handlers package
	app.Rabbit = rabbitConn
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Start web server
	if err := web.Run(multiplexer(), webPort); err != nil {
		log.Fatalf(err.Error())
	}

}
