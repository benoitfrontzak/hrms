package main

import (
	"broker/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// Statics routes
	mux.HandleFunc("/sendEmail", handlers.Repo.SendEmail)
	mux.HandleFunc("/createEmployee", handlers.Repo.CreateEmployee)

	return mux
}
