package handlers

import (
	"leave/pkg/pg"
	"log"
	"net/http"
	"time"
)

// CreateMyLeave is the handler which receives
// a leave payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) CreateMyLeave(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p pg.Leave

	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	rep.App.Models.Leave = p
	log.Println("payload received", p)
	// insert payload to employee's config table DB
	rowID, err := rep.App.Models.Leave.Insert()
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
