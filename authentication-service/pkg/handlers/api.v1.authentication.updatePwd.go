package handlers

import (
	"authentication/pkg/pg"
	"errors"
	"log"
	"net/http"
)

// Update pwd in authentication-service
// Update pwd in authentication-service
func (rep *Repository) UpdatePwd(w http.ResponseWriter, r *http.Request) {
	// extract payload from request
	var pw pg.Pwd
	err := rep.readJSON(w, r, &pw)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	log.Println("payload", pw)

	//validate and convert payload to user to be reuse
	u, err := rep.validatePassword(pw)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// insert user to DB
	err = u.ResetPassword(pw.NewPassword)
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

// func comparePasswords(hashedPassword string, plainPassword string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
// 	return err == nil
// }

// convert payload to user
func (rep *Repository) validatePassword(payload pg.Pwd) (pg.User, error) {
	var u pg.User
	log.Println("id", payload.ID)
	u.ID = payload.ID
	dbPw, err := rep.App.Models.User.GetPassword(payload.ID)
	if err != nil {
		return u, err
	}
	u.Password = dbPw
	check, err := u.PasswordMatches(payload.OldPassword)
	if err != nil {
		return u, err
	}
	if !check {
		return u, errors.New("passwords do not match")
	}

	return u, nil
}
