package main

import (
	"authentication/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// Statics route public
	mux.HandleFunc("/api/v1/authentication/authenticate", handlers.Repo.Authenticate)
	mux.HandleFunc("/api/v1/authentication/verifyCookie", handlers.Repo.VerifyCookie)
	mux.HandleFunc("/api/v1/authentication/allUsers", handlers.Repo.AllUsers)
	mux.HandleFunc("/api/v1/authentication/createUser", handlers.Repo.CreateUser)
	mux.HandleFunc("/api/v1/authentication/updateUser", handlers.Repo.UpdateUser)
	mux.HandleFunc("/api/v1/authentication/update/password", handlers.Repo.UpdatePassword)

	return mux
}
