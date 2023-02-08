package main

import (
	"embed"
	"frontend/pkg/handlers"
	"net/http"
)

//go:embed static
var staticFiles embed.FS

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// Files server (Static files)
	mux.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	// Statics route public
	mux.HandleFunc("/", handlers.Repo.Login)
	mux.HandleFunc("/unauthorized", handlers.Repo.Unauthorized)
	// mux.HandleFunc("/home", handlers.Repo.Home)
	mux.HandleFunc("/logout", handlers.Repo.Logout)

	// With middleware
	mux.Handle("/home", handlers.Middleware(http.HandlerFunc(handlers.Repo.Home)))
	mux.Handle("/mailer", handlers.Middleware(http.HandlerFunc(handlers.Repo.Mailer)))
	mux.Handle("/demo", handlers.Middleware(http.HandlerFunc(handlers.Repo.Demo)))

	return mux
}
