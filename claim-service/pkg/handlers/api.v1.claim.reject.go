package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

// ApproveClaimCT is the handler which receives
// an claim payload from the broker
// convert it and updates it to the DB
func (rep *Repository) RejectClaim(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var p listID
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	log.Println("payload from request is:", p)
	// approve each claim to DB
	for _, rid := range p.List {
		rowID, _ := strconv.Atoi(rid)
		err := rep.App.Models.Claim.Reject(rowID, p.UserID, p.Reason)
		if err != nil {
			log.Println("err db is:", err)
			rep.errorJSON(w, err)
			return
		}
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "claims successfully approved",
		Data:      p,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p.ConnectedUser,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
