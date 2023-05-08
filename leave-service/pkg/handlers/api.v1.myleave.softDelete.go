package handlers

import (
	"net/http"
	"strconv"
	"time"
)

// receives a get request from broker to fetch all employee
func (rep *Repository) SoftDeleteMyLeave(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var p listID

	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// soft delete my leaves one by one
	for _, eid := range p.List {
		id, _ := strconv.Atoi(eid)
		// soft delete leave application
		err = rep.App.Models.Leave.Delete(id, p.UserID)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}
		// get leave definition id of leave application id
		lid, err := rep.App.Models.Leave.GetLeaveDefinitionID(id)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}
		// calculate number of days to be removed to taken
		n, err := rep.App.Models.Leave.GetRequestedDatesNumber(id)
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
		Message:   "leave successfully deleted",
		Data:      p.List,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p.ConnectedUser,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
