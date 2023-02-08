package main

import (
	"authentication/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// Statics route public
	mux.HandleFunc("/authenticate", handlers.Repo.Authenticate)
	mux.HandleFunc("/verifyCookie", handlers.Repo.VerifyCookie)
	mux.HandleFunc("/allUsers", handlers.Repo.AllUsers)

	return mux
}
