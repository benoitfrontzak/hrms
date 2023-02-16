package main

import (
	"employee/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// Statics route public
	mux.HandleFunc("/createEmployee", handlers.Repo.CreateEmployee)

	return mux
}
