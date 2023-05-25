package handlers

import (
	"employee/pkg/pg"
	"net/http"
	"time"
)

// UpdatePI is the handler which receives
// an payroll item payload from the broker
// convert it and update it to the DB
func (rep *Repository) UpdatePI(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p pg.PayrollItem
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// insert payload to employee DB
	err = rep.App.Models.Employee.UpdatePI(p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "payroll item successfully updated",
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
