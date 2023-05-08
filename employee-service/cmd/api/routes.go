package main

import (
	"employee/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// EMPLOYEE CUD
	mux.HandleFunc("/api/v1/employee/create", handlers.Repo.Create)         // create new employee
	mux.HandleFunc("/api/v1/employee/update", handlers.Repo.Update)         // update employee
	mux.HandleFunc("/api/v1/employee/softDelete", handlers.Repo.SoftDelete) // delete employee

	// EMPLOYEE GET
	mux.HandleFunc("/api/v1/employee/get/seniority", handlers.Repo.GetSeniorityByID) // return seniority of employee id (year int)
	mux.HandleFunc("/api/v1/employee/get/all", handlers.Repo.GetAll)                 // return all employee by status (active, inactive, deleted)
	mux.HandleFunc("/api/v1/employee/get/allActive", handlers.Repo.GetAllActive)     // return list of all active employees
	mux.HandleFunc("/api/v1/employee/get/id", handlers.Repo.GetByID)                 // return employee information by id
	mux.HandleFunc("/api/v1/employee/get/email", handlers.Repo.GetByEmail)           // return employee information by email
	mux.HandleFunc("/api/v1/employee/get/leave", handlers.Repo.GetLeaveInfo)         // return employee information relative to leave-service
	mux.HandleFunc("/api/v1/employee/find/email", handlers.Repo.FindEmail)           // return employee information by email

	// CONFIG TABLES CRUD
	mux.HandleFunc("/api/v1/employee/configTables/get", handlers.Repo.GetCT)              // return all employee's config table
	mux.HandleFunc("/api/v1/employee/configTable/create", handlers.Repo.CreateCT)         // create new config table entry
	mux.HandleFunc("/api/v1/employee/configTable/update", handlers.Repo.UpdateCT)         // update one config table entry
	mux.HandleFunc("/api/v1/employee/configTable/softDelete", handlers.Repo.SoftDeleteCT) // delete list of config table entries

	return mux
}
