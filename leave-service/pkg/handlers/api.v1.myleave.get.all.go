package handlers

import (
	"log"
	"net/http"
	"time"
)

// GetAllMyLeave is the handler which receives a payload from the broker
// and fetch all my leaves by employee ID
func (rep *Repository) GetAllMyLeave(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllMyLeave")
	// extract payload from request
	var p User
	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println("json err", err)
		rep.errorJSON(w, err)
		return
	}

	log.Println("payload received:", p)

	// get all my leaves
	all, err := rep.App.Models.Leave.GetAllMyLeave(p.ID)
	if err != nil {
		log.Println("db err:", err)
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
	log.Println("answer:", answer)
	rep.writeJSON(w, http.StatusAccepted, answer)

}
