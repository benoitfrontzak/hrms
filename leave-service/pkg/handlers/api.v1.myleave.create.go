package handlers

import (
	"leave/pkg/pg"
	"net/http"
	"time"
)

// CreateMyLeave is the handler which receives
// a leave payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) CreateMyLeave(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p pg.Leave

	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	rep.App.Models.Leave = p

	// insert payload to employee's config table DB
	rowID, err := rep.App.Models.Leave.Insert()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	// calculate number of days to be added to taken
	n, err := rep.App.Models.Leave.GetRequestedDatesNumber(rowID)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// increment employee's leave taken
	err = rep.App.Models.Leave.IncrementTakenLeave(p.LeaveDefinitionID, p.EmployeeID, n)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "new config table entry successfully created",
		Data:      rowID,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: "p.ConnectedUser",
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
