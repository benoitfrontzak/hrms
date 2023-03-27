package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

// ApproveLeave is the handler which approve leave application
func (rep *Repository) ApproveLeave(w http.ResponseWriter, r *http.Request) {
	log.Println("ApproveLeave hit")
	// extract payload from request
	var p listID
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// approve each claim to DB
	for _, rid := range p.List {
		rowID, _ := strconv.Atoi(rid)
		err := rep.App.Models.Leave.Approve(rowID, p.UserID)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "leaves successfully approved",
		Data:      p,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p.ConnectedUser,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
