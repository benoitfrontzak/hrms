package handlers

import (
	"leave/pkg/pg"
	"net/http"
	"time"
)

// CreateLeaveDefinition is the handler which receives
// a leave definition payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) CreateLeaveDefinition(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p pg.LeaveDefinition

	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// insert payload to employee's config table DB
	rep.App.Models.LeaveDefinition = p
	_, err = rep.App.Models.LeaveDefinition.Insert()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "new leave definition successfully created",
		Data:      p,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: "p.ConnectedUser",
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
