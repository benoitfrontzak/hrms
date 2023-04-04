package main

import (
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// Statics route public
	// todo

	return mux
}
