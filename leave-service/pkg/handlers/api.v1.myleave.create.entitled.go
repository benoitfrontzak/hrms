package handlers

import (
	"log"
	"net/http"
	"time"
)

// CreateMyEntitledLeave is the handler which receives
// an employee id payload from the broker
// convert it and creates all entitled leave to LEAVE_EMPLOYEE
func (rep *Repository) CreateMyEntitledLeave(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p User

	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println("json decoding err: ", err)
		rep.errorJSON(w, err)
		return
	}

	log.Println("payload received", p)

	// generate all entitled leaves for employe_id to LEAVE_EMPLOYEE
	err = rep.App.Models.Leave.CreateEntitled(p.ID, p.UserID, p.GenderID)
	if err != nil {
		log.Println("db err: ", err)
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "entitled leaves successfully created for employee " + p.Email,
		Data:      "",
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p.ConnectedUser,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
