package handlers

import (
	"net/http"
	"time"
)

// RejectLeave is the handler which reject leave application
func (rep *Repository) UpdateCredits(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var p []ListCredits
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// update DB
	for _, element := range p {
		err := rep.App.Models.Leave.UpdateCredits(element.RowID, element.UserID, element.Credits)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "leave successfully rejected",
		Data:      p,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p[0].ConnectedUser,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
