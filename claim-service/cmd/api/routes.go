package main

import (
	"claim/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// API v1 Claim Definition
	mux.HandleFunc("/api/v1/claim/definition/get/all", handlers.Repo.GetAllClaimDefinition)
	mux.HandleFunc("/api/v1/claim/definition/create", handlers.Repo.CreateClaimDefinition)
	mux.HandleFunc("/api/v1/claim/definition/update", handlers.Repo.UpdateClaimDefinition)
	mux.HandleFunc("/api/v1/claim/definition/softDelete", handlers.Repo.SoftDeleteClaimDefinition)

	// API v1 Config tables
	mux.HandleFunc("/api/v1/claim/configTables/get", handlers.Repo.GetAllConfigTablesClaim)
	mux.HandleFunc("/api/v1/claim/configTable/create", handlers.Repo.CreateClaimCT)
	mux.HandleFunc("/api/v1/claim/configTable/update", handlers.Repo.UpdateClaimCT)
	mux.HandleFunc("/api/v1/claim/configTable/softDelete", handlers.Repo.SoftDeleteClaimCT)

	// API v1 My Claim
	mux.HandleFunc("/api/v1/myclaim/get/all", handlers.Repo.GetAllMyClaim)
	mux.HandleFunc("/api/v1/myclaim/create", handlers.Repo.CreateMyClaim)
	// mux.HandleFunc("/api/v1/myclaim/update", handlers.Repo.UpdateMyClaim)
	mux.HandleFunc("/api/v1/myclaim/softDelete", handlers.Repo.SoftDeleteMyClaim)

	// API v1 claim
	mux.HandleFunc("/api/v1/claim/get/all", handlers.Repo.GetAllClaim)
	mux.HandleFunc("/api/v1/claim/approve", handlers.Repo.ApproveClaim)
	mux.HandleFunc("/api/v1/claim/reject", handlers.Repo.RejectClaim)

	return mux
}
