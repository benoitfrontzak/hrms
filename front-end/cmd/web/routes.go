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

	// Files server (Static files embed)
	mux.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	// File server (files that can be downloaded dynamicaly)
	uploadedFiles := "/upload/"
	fs := http.FileServer(http.Dir("." + uploadedFiles))
	mux.Handle(uploadedFiles, http.StripPrefix(uploadedFiles, fs))

	// Public
	mux.HandleFunc("/", handlers.Repo.Login)
	mux.HandleFunc("/unauthorized", handlers.Repo.Unauthorized)
	mux.HandleFunc("/logout", handlers.Repo.Logout)

	// Home
	mux.Handle("/home", handlers.Middleware(http.HandlerFunc(handlers.Repo.Home))) // Dashboard

	// Information: super admin only (logs analysis)
	mux.Handle("/information/", handlers.Middleware(http.HandlerFunc(handlers.Repo.Information)))

	// Maintenance: Admin only
	mux.Handle("/employee/", handlers.Middleware(http.HandlerFunc(handlers.Repo.Employee))) // CRUD operation
	mux.Handle("/claim/", handlers.Middleware(http.HandlerFunc(handlers.Repo.Claim)))       // Approve | reject claim
	mux.Handle("/leave/", handlers.Middleware(http.HandlerFunc(handlers.Repo.Leave)))       // Approve | reject leave

	// My profile
	mux.Handle("/myclaim/", handlers.Middleware(http.HandlerFunc(handlers.Repo.MyClaim))) // my claim applications (CRUD)
	mux.Handle("/myleave/", handlers.Middleware(http.HandlerFunc(handlers.Repo.MyLeave))) // my leave applications (CRUD)

	// configuration
	mux.Handle("/configuration/employee/", handlers.Middleware(http.HandlerFunc(handlers.Repo.EmployeeConfiguration)))
	mux.Handle("/configuration/claim/", handlers.Middleware(http.HandlerFunc(handlers.Repo.ClaimConfiguration)))
	mux.Handle("/configuration/leave/", handlers.Middleware(http.HandlerFunc(handlers.Repo.LeaveConfiguration)))

	// Definition
	mux.Handle("/definition/claim/", handlers.Middleware(http.HandlerFunc(handlers.Repo.ClaimDefinition)))
	mux.Handle("/definition/leave/", handlers.Middleware(http.HandlerFunc(handlers.Repo.LeaveDefinition)))

	// API with middleware (upload)
	mux.Handle("/api/v1/employee/upload/", handlers.Middleware(http.HandlerFunc(handlers.Repo.UploadFiles)))
	mux.Handle("/api/v1/employee/getUploadedFiles/", handlers.Middleware(http.HandlerFunc(handlers.Repo.GetUploadedFiles)))

	return mux
}
