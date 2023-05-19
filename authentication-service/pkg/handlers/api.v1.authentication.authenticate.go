package handlers

import (
	"authentication/pkg/token"
	"errors"
	"net/http"
)

// Authenticate checks if user: login/password is matching a db record
// if successful, it creates a paseto token and encode it in a secure httponly cookie
// if failed, it redirects to unauthorized page
// it sends the access attempt (successful or failed) to the logger service
func (rep *Repository) Authenticate(w http.ResponseWriter, r *http.Request) {
	// Invoke ParseForm before read form values
	r.ParseForm()
	// Extract payload from request
	requestPayload := AuthPayload{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// Validate the user against the database
	user, err := rep.App.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		http.Redirect(w, r, frontEnd+"unauthorized", http.StatusSeeOther)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		http.Redirect(w, r, frontEnd+"unauthorized", http.StatusSeeOther)
		return
	}

	// User is authenticated so let's create the access token (paseto)
	accessToken, err := rep.App.Paseto.CreateToken(user.Email, user.Nickname, user.Role, user.EmployeeID)
	if err != nil {
		rep.errorJSON(w, errors.New("cannot create paseto token"), http.StatusUnauthorized)
		return
	}

	// Encode token into a secure cookie
	token.SetCookie(accessToken, w)

	// All set, redirect to home page with secure cookie in response header
	http.Redirect(w, r, frontEnd+"home", http.StatusSeeOther)

}
