package handlers

import (
	"net/http"
)

// GetAllConfigTablesClaim is the handler which returns all claim's CT
func (rep *Repository) GetAllConfigTablesClaim(w http.ResponseWriter, r *http.Request) {
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
