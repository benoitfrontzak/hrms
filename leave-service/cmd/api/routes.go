package main

import (
	"leave/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// API v1 Leave Definition
	mux.HandleFunc("/api/v1/leave/definition/get/all", handlers.Repo.GetAllLeaveDefinition)
	mux.HandleFunc("/api/v1/leave/definition/create", handlers.Repo.CreateLeaveDefinition)
	mux.HandleFunc("/api/v1/leave/definition/update", handlers.Repo.UpdateLeaveDefinition)
	mux.HandleFunc("/api/v1/leave/definition/softDelete", handlers.Repo.SoftDeleteLeaveDefinition)

	// API v1 Config tables
	mux.HandleFunc("/api/v1/leave/configTables/get", handlers.Repo.GetAllConfigTablesLeave)
	mux.HandleFunc("/api/v1/leave/configTable/create", handlers.Repo.CreateLeaveCT)
	mux.HandleFunc("/api/v1/leave/configTable/update", handlers.Repo.UpdateLeaveCT)
	mux.HandleFunc("/api/v1/leave/configTable/softDelete", handlers.Repo.SoftDeleteLeaveCT)

	// API v1 My Leave
	mux.HandleFunc("/api/v1/myleave/get/all", handlers.Repo.GetAllMyLeave)
	mux.HandleFunc("/api/v1/myleave/create", handlers.Repo.CreateMyLeave)
	// mux.HandleFunc("/api/v1/myleave/update", handlers.Repo.UpdateMyLeave)
	mux.HandleFunc("/api/v1/myleave/softDelete", handlers.Repo.SoftDeleteMyLeave)

	// API v1 leave
	mux.HandleFunc("/api/v1/leave/get/all", handlers.Repo.GetAllLeave)
	mux.HandleFunc("/api/v1/leave/approve", handlers.Repo.ApproveLeave)
	mux.HandleFunc("/api/v1/leave/reject", handlers.Repo.RejectLeave)

	return mux
}
