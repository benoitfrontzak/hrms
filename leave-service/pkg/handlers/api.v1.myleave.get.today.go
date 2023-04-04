package handlers

import (
	"log"
	"net/http"
	"time"
)

// GetAllMyLeaveToday is the handler which receives a payload from the broker
// and fetch all my entitled leaves by employee ID
func (rep *Repository) GetAllMyLeaveToday(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var p User
	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println("json err", err)
		rep.errorJSON(w, err)
		return
	}

	// get all my entitled leaves
	all, err := rep.App.Models.Leave.GetAllMyEntitledLeave(p.ID)
	if err != nil {
		log.Println("db err", err)
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "my leave successfully fetched",
		Data:      all,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p.Email,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
