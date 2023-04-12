package handlers

import (
	"net/http"
)

// GetAllLeaveDefinition is the handler which returns all claim definition
func (rep *Repository) GetLeaveDefinitionID(w http.ResponseWriter, r *http.Request) {
	// extract leave definition id
	type payload struct {
		ID int
	}
	var p payload
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// get leave definition id information
	all, err := rep.App.Models.LeaveDefinition.GetLeaveDefinitionID(p.ID)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:   false,
		Message: "claim definition successfully fetched",
		Data:    all,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
