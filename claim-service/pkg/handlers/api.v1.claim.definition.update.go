package handlers

import (
	"claim/pkg/pg"
	"log"
	"net/http"
	"time"
)

// UpdateClaimDefinition is the handler which receives
// a payload from the broker and updates it to the DB
func (rep *Repository) UpdateClaimDefinition(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p pg.ClaimDefinition

	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println("error in UPdate readjson", err)
		rep.errorJSON(w, err)
		return
	}

	// insert payload to employee's config table DB
	rep.App.Models.ClaimDefinition = p
	err = rep.App.Models.ClaimDefinition.Update()
	if err != nil {
		log.Println("error at update", err)
		rep.errorJSON(w, err)
		return
	}

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "new config table entry successfully created",
		Data:      p,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: "p.ConnectedUser",
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
