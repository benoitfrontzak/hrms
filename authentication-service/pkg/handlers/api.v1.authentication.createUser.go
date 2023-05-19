package handlers

import (
	"authentication/pkg/pg"
	"net/http"
	"time"
)

// CreateUser creates a new user to authentication-service
func (rep *Repository) CreateUser(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var u pg.User
	err := rep.readJSON(w, r, &u)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// insert user to DB
	rowID, err := rep.App.Models.User.Insert(u)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// send response
	resp := jsonResponse{
		Error:     false,
		Message:   "user successfully created",
		Data:      rowID,
		CreatedAt: time.Now().Format("02-Jan-2006 15:04:05"),
		CreatedBy: "",
	}
	noToken := ""
	rep.writeJSON(w, http.StatusAccepted, resp, noToken)

}
