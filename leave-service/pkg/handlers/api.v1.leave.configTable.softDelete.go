package handlers

import (
	"errors"
	"net/http"
	"time"
)

// SoftDeleteLeaveCT is the handler which receives
// an leave's CT payload from the broker
// convert it and inserts it to the DB
func (rep *Repository) SoftDeleteLeaveCT(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var p ctEntriesPayload
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

	// sof delete each config table entry
	for _, id := range p.List {
		err := rep.App.Models.ConfigTables.Delete(dbTableName, id)
		if err != nil {
			rep.errorJSON(w, err)
			return
		}
	}

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
