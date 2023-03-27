package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"
)

// UpdateLeaveCT is the handler which receives
// an leave's CT payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) UpdateLeaveCT(w http.ResponseWriter, r *http.Request) {

	// extract payload from request
	var p configTablePayload
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// get db table name
	dbTableName := getTableNameDB(p.Table)
	if dbTableName == "" {
		rep.errorJSON(w, errors.New("no such config table in the db"))
		return
	}

	// update payload to employee's config table DB
	rowID, err := rep.App.Models.ConfigTables.Update(dbTableName, p.Entry, p.RowID)
	if err != nil {
		log.Println("err db is:", err)
		rep.errorJSON(w, err)
		return
	}
	p.RowID = rowID

	// response to be sent
	answer := jsonResponse{
		Error:     false,
		Message:   "new config table entry successfully created",
		Data:      p,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: p.ConnectedUser,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)

}
