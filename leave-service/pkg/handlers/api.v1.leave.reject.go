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
		// get leave definition id of leave application id
		lid, err := rep.App.Models.Leave.GetLeaveDefinitionID(rowID)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}
		// calculate number of days to be removed to taken
		n, err := rep.App.Models.Leave.GetRequestedDatesNumber(rowID)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}

		// decrement employee's leave taken
		err = rep.App.Models.Leave.DecrementTakenLeave(lid, p.UserID, n)
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
