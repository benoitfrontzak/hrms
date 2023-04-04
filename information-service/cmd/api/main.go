package main

import (
	"context"
	"information/pkg/config"
	"information/pkg/handlers"
	"information/pkg/mongoDB"
	"information/pkg/web"
	"log"
	"time"

	_ "time/tzdata"
)

const (
	webPort = "8181"
)

func main() {
	// Customize log display
	setLog()

	// Set time zone
	tz, err := time.LoadLocation("Asia/Kuala_Lumpur")
	if err != nil {
		log.Println("error while setting the time zone: ", err)
	}
	// this is setting the global timezone
	time.Local = tz

	// connect to mongo
	conn, err := mongoDB.ConnectToDB()
	if err != nil {
		log.Panic(err)
	}

	// Set connexion (*mongo.Client) for pkg mongoDB
	mongoDB.Client = conn

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = conn.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// app holds our application configuration
	app := config.AppConfig{
		DB:     conn,
		Models: mongoDB.New(conn),
	}

	// set the config for the handlers package
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Start web server
	if err := web.Run(multiplexer(), webPort); err != nil {
		log.Fatalf(err.Error())
	}
}
