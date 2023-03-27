package handlers

import (
	"net/http"
)

// CreateEmployee is the handler which receives
// an employee payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) GetCT(w http.ResponseWriter, r *http.Request) {
	// get all employee's config tables
	ctables, err := rep.App.Models.ConfigTables.GetAllConfigTables()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:   false,
		Message: "employee CT successfully fetched",
		Data:    ctables,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
