package handlers

import (
	"leave/pkg/pg"
	"net/http"
	"time"
)

// UpdateClaimDefinition is the handler which receives
// a payload from the broker and updates it to the DB
func (rep *Repository) UpdateLeaveDefinition(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p pg.LeaveDefinition
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// insert payload to employee's config table DB
	rep.App.Models.LeaveDefinition = p
	err = rep.App.Models.LeaveDefinition.Update()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "leave definition successfully updated",
		Data:      p,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: "p.ConnectedUser",
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
