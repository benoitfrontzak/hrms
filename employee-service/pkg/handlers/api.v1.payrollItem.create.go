package handlers

import (
	"employee/pkg/pg"
	"net/http"
	"time"
)

// CreateNewPI is the handler which receives
// an payroll item payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) CreateNewPI(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p pg.PayrollItem
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// insert payload to employee DB
	rowID, err := rep.App.Models.Employee.InsertNewPI(p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "payroll item successfully created",
		Data:      rowID,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
