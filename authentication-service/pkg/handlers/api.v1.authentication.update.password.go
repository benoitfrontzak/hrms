package handlers

import (
	"authentication/pkg/pg"
	"errors"
	"log"
	"net/http"
)

// Update password for one employee ID
func (rep *Repository) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var payload pg.ChangePassword
	err := rep.readJSON(w, r, &payload)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	log.Println(payload)
	// fetch user information from DB
	user, err := rep.App.Models.User.GetOne(payload.CreatedBy)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	log.Println(user.Fullname)
	log.Println(user.Email)
	log.Println(user.ID)
	log.Println(user.EmployeeID)
	//validate and convert payload to user to be reuse
	valid, err := user.PasswordMatches(payload.OldPassword)
	log.Println(valid)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	if !valid {
		rep.errorJSON(w, errors.New("old password is not matching"))
		return
	}

	user.Password = payload.NewPassword
	user.EmployeeID = payload.CreatedBy

	// insert user to DB
	err = user.UpdatePassword()
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// send response
	resp := jsonResponse{
		Error:   false,
		Message: "user password successfully updated",
		Data:    nil,
	}
	noToken := ""
	rep.writeJSON(w, http.StatusAccepted, resp, noToken)
}
