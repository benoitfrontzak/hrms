package handlers

import (
	"log"
	"net/http"
)

// GetAllLeave is the handler which returns all leave definition
func (rep *Repository) GetAllLeaveUser(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var p User

	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println("json err", err)
		rep.errorJSON(w, err)
		return
	}

	log.Println("payload", p)
	// get leaves (approved, rejected, pending)
	all, err := rep.App.Models.Leave.GetAllLeaveUser(p.ID)
	if err != nil {
		log.Println("db err", err)
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
