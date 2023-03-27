package handlers

import (
	"claim/pkg/pg"
	"log"
	"net/http"
	"time"
)

// CreateClaimDefinition is the handler which receives
// a claim definition payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) CreateClaimDefinition(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p pg.ClaimDefinition

	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println("pd is err", err)
		rep.errorJSON(w, err)
		return
	}

	// insert payload to employee's config table DB
	rep.App.Models.ClaimDefinition = p
	_, err = rep.App.Models.ClaimDefinition.Insert()
	if err != nil {
		log.Println("err db is:", err)
		rep.errorJSON(w, err)
		return
	}
	// p.RowID = rowID

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
