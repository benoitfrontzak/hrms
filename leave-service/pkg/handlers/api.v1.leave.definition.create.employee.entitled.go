package handlers

import (
	"net/http"
	"time"
)

// UpdateClaimDefinition is the handler which receives
// a payload from the broker and updates it to the DB
func (rep *Repository) CreateEntitledByDefinition(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	type payload struct {
		LeaveID    int
		EmployeeID int
		Entitled   float64
	}
	var p payload
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// insert payload to employee's config table DB
	err = rep.App.Models.LeaveDefinition.CreateEntitledByDefinition(p.LeaveID, p.EmployeeID, p.Entitled)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	var empty any
	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "leave employee new entitlement successfully created",
		Data:      empty,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: "p.ConnectedUser",
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
