package handlers

import (
	"net/http"
)

// GetAllLeave is the handler which returns all leave definition
func (rep *Repository) GetAllLeave(w http.ResponseWriter, r *http.Request) {
	// get leaves (approved, rejected, pending)
	all, err := rep.App.Models.Leave.GetAllLeave()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:   false,
		Message: "leave successfully fetched",
		Data:    all,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
