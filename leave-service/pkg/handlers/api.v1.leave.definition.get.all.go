package handlers

import (
	"net/http"
)

// GetAllLeaveDefinition is the handler which returns all claim definition
func (rep *Repository) GetAllLeaveDefinition(w http.ResponseWriter, r *http.Request) {
	// get claim definition (active, inactive, deleted)
	all, err := rep.App.Models.LeaveDefinition.GetLeaveDefinition()
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
