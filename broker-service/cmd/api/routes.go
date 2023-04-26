package main

import (
	"broker/pkg/handlers"
	"net/http"
)

// multiplexer is the http.Handler for broker-service
func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// authentication-service routes
	mux.HandleFunc("/route/authentication/update/password", handlers.UpdatePassword) // update employee's password
	// employee-service routes
	mux.HandleFunc("/route/employee/get/all", handlers.AllEmployees)        // return all employee by status (active, inactive, deleted)
	mux.HandleFunc("/route/employee/get/id", handlers.EmployeeByID)         // return employee information by id
	mux.HandleFunc("/route/employee/get/email", handlers.EmployeeByEmail)   // return employee information by email
	mux.HandleFunc("/route/employee/get/leave", handlers.EmployeeLeaveInfo) // return employee information relative to leave-service

	mux.HandleFunc("/route/employee/create", handlers.CreateEmployee)         // create new employee
	mux.HandleFunc("/route/employee/update", handlers.UpdateEmployee)         // update employee
	mux.HandleFunc("/route/employee/softDelete", handlers.SoftDeleteEmployee) // delete employee

	mux.HandleFunc("/route/employee/configTables/get", handlers.EmployeeCT)                 // return all employee's config table
	mux.HandleFunc("/route/employee/configTable/create", handlers.CreateEmployeeCT)         // create new config table entry
	mux.HandleFunc("/route/employee/configTable/update", handlers.UpdateEmployeeCT)         // update one config table entry
	mux.HandleFunc("/route/employee/configTable/softDelete", handlers.SoftDeleteEmployeeCT) // delete list of config table entries

	// claim-service routes
	mux.HandleFunc("/route/myclaim/get/all", handlers.AllMyClaim)                          // return all my claims
	mux.HandleFunc("/route/myclaim/get/thisYear", handlers.AllMyYearlyClaim)               // return all my claims of the year (total)
	mux.HandleFunc("/route/myclaim/get/thisYearDetails", handlers.AllMyYearlyClaimDetails) // return all my claims of the year (details)
	mux.HandleFunc("/route/myclaim/create", handlers.CreateMyClaim)                        // create my claim
	// mux.HandleFunc("/route/myclaim/update", handlers.UpdateMyClaim)         // update my claim
	mux.HandleFunc("/route/myclaim/softDelete", handlers.SoftDeleteMyClaim) // delete my claim

	mux.HandleFunc("/route/claim/get/all", handlers.AllClaim)     // return all claims
	mux.HandleFunc("/route/claim/approve", handlers.ApproveClaim) // approve all claims
	mux.HandleFunc("/route/claim/reject", handlers.RejectClaim)   // reject all claims

	mux.HandleFunc("/route/claim/definition/get/all", handlers.AllClaimDefinition)           // return all claim definitions
	mux.HandleFunc("/route/claim/definition/create", handlers.CreateClaimDefinition)         // create new claim definition
	mux.HandleFunc("/route/claim/definition/update", handlers.UpdateClaimDefinition)         // update one claim definition
	mux.HandleFunc("/route/claim/definition/softDelete", handlers.SoftDeleteClaimDefinition) // delete one claim definition

	mux.HandleFunc("/route/claim/configTables/get", handlers.ClaimCT)                          // return all claim's config table
	mux.HandleFunc("/route/claim/configTable/create", handlers.ClaimConfigTableCreate)         // create new config table entry
	mux.HandleFunc("/route/claim/configTable/update", handlers.ClaimConfigTableUpdate)         // update one config table entry
	mux.HandleFunc("/route/claim/configTable/softDelete", handlers.ClaimConfigTableSoftDelete) // delete list of config table entries

	// leave-service routes
	mux.HandleFunc("/route/myleave/get/all", handlers.AllMyLeave)        // return all my leaves
	mux.HandleFunc("/route/myleave/get/today", handlers.AllMyLeaveToday) // return all my leaves of the year (details)
	mux.HandleFunc("/route/myleave/create", handlers.CreateMyLeave)      // create my leave
	// mux.HandleFunc("/route/myleave/update", handlers.UpdateMyLeave)         // update my leave
	mux.HandleFunc("/route/myleave/softDelete", handlers.SoftDeleteMyLeave) // delete my leave

	mux.HandleFunc("/route/leave/get/all", handlers.AllLeave)        // return all leaves
	mux.HandleFunc("/route/leave/get/today", handlers.AllLeaveToday) // return all leaves of today & tomorrow
	mux.HandleFunc("/route/leave/approve", handlers.ApproveLeave)    // approve all leaves
	mux.HandleFunc("/route/leave/reject", handlers.RejectLeave)      // reject all leaves

	mux.HandleFunc("/route/leave/definition/get/all", handlers.AllLeaveDefinition)           // return all leave definitions
	mux.HandleFunc("/route/leave/definition/create", handlers.CreateLeaveDefinition)         // create new leave definition
	mux.HandleFunc("/route/leave/definition/update", handlers.UpdateLeaveDefinition)         // update one leave definition
	mux.HandleFunc("/route/leave/definition/softDelete", handlers.SoftDeleteLeaveDefinition) // delete one leave definition

	mux.HandleFunc("/route/leave/configTables/get", handlers.LeaveCT)                          // return all leave's config table
	mux.HandleFunc("/route/leave/configTable/create", handlers.LeaveConfigTableCreate)         // create new config table entry
	mux.HandleFunc("/route/leave/configTable/update", handlers.LeaveConfigTableUpdate)         // update one config table entry
	mux.HandleFunc("/route/leave/configTable/softDelete", handlers.LeaveConfigTableSoftDelete) // delete list of config table entries
	return mux
}
