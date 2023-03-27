package handlers

import (
	"net/http"
	"strconv"
	"time"
)

// RejectLeave is the handler which reject leave application
func (rep *Repository) RejectLeave(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var p listID
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// update DB
	for _, rid := range p.List {
		rowID, _ := strconv.Atoi(rid)
		err := rep.App.Models.Leave.Reject(rowID, p.UserID, p.Reason)
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
		CreatedBy: p.ConnectedUser,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
