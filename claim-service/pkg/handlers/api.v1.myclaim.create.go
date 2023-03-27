package handlers

import (
	"claim/pkg/pg"
	"net/http"
	"time"
)

// CreateMyClaim is the handler which receives
// a claim payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) CreateMyClaim(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p pg.Claim

	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// insert payload to employee's config table DB
	rep.App.Models.Claim = p
	rowID, err := rep.App.Models.Claim.Insert()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	// p.RowID = rowID

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "new config table entry successfully created",
		Data:      rowID,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: "p.ConnectedUser",
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
