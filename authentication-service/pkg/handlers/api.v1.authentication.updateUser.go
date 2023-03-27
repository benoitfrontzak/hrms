package handlers

import (
	"authentication/pkg/pg"
	"net/http"
)

// UpdateUser updates a user to authentication-service
func (rep *Repository) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var u pg.User
	err := rep.readJSON(w, r, &u)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// insert user to DB
	err = rep.App.Models.User.Update(u)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// send response
	resp := jsonResponse{
		Error:   false,
		Message: "user successfully updated",
		Data:    nil,
	}
	noToken := ""
	rep.writeJSON(w, http.StatusAccepted, resp, noToken)

}
