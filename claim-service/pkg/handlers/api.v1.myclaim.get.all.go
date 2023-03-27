package handlers

import (
	"log"
	"net/http"
	"time"
)

// GetAllMyClaim is the handler which receives a payload from the broker
// and fetch all my claims by employee ID
func (rep *Repository) GetAllMyClaim(w http.ResponseWriter, r *http.Request) {
	log.Println("claim hit")
	// extract payload from request
	var p User
	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println("decoded json:", err)
		rep.errorJSON(w, err)
		return
	}

	log.Println("employee requested ID:", p)

	// get all my claims
	all, err := rep.App.Models.Claim.GetAllMyClaim(p.ID)
	if err != nil {
		log.Println("db err:", err)
		rep.errorJSON(w, err)
		return
	}

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
