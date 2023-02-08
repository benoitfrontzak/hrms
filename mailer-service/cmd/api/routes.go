package main

import (
	"mailer/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// Statics route public
	mux.HandleFunc("/sendEmail", handlers.Repo.Mailer)

	return mux
}
