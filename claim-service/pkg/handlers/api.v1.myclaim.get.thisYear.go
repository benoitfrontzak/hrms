package handlers

import (
	"log"
	"net/http"
	"time"
)

// AllMyYearlyClaim is the handler which receives a payload from the broker
// and fetch all my claims of the year by employee ID
func (rep *Repository) AllMyYearlyClaim(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p User
	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println("err json")
		log.Println(err)
		rep.errorJSON(w, err)
		return
	}

	// get all my claims
	all, err := rep.App.Models.Claim.GetAllMyYearlyClaim(p.ID)
	if err != nil {
		log.Println("err db")
		log.Println(err)
		rep.errorJSON(w, err)
		return
	}

	log.Println("all")
	log.Println(all)
	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "my claims successfully fetched",
		Data:      all,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p.Email,
	}
	log.Println("answer:", answer)
	rep.writeJSON(w, http.StatusAccepted, answer)

}
