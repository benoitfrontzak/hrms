package handlers

import (
	"authentication/pkg/token"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
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
		Email:    r.FormValue("email_input"),
		Password: r.FormValue("password_input"),
	}

	// Validate the user against the database
	user, err := rep.App.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		// Log invalid user credentials authentication access to logger service
		rpcPayload := RPCPayload{
			Collection: "authentication",
			Name:       "Authentication failed",
			Data:       fmt.Sprintf("invalid user credentials for user: %s at %s", requestPayload.Email, time.Now().Format("2006-01-02 15:04:05")),
		}
		rep.LogItemViaRPC(w, rpcPayload)
		log.Printf("invalid user credentials for user: %s at %s", requestPayload.Email, time.Now().Format("2006-01-02 15:04:05"))
		http.Redirect(w, r, "http://localhost/unauthorized", http.StatusSeeOther)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		// Log invalid user credentials authentication access to logger service
		rpcPayload := RPCPayload{
			Collection: "authentication",
			Name:       "Authentication failed",
			Data:       fmt.Sprintf("invalid password credentials for user: %s at %s", requestPayload.Email, time.Now().Format("2006-01-02 15:04:05")),
		}
		rep.LogItemViaRPC(w, rpcPayload)
		log.Printf("invalid password credentials for user: %s at %s", requestPayload.Email, time.Now().Format("2006-01-02 15:04:05"))
		http.Redirect(w, r, "http://localhost/unauthorized", http.StatusSeeOther)
		return
	}

	// Log successful authentication access to logger service
	rpcPayload := RPCPayload{
		Collection: "authentication",
		Name:       "Authentication successful",
		Data:       fmt.Sprintf("Logged in user %s at %s", user.Email, time.Now().Format("2006-01-02 15:04:05")),
	}
	rep.LogItemViaRPC(w, rpcPayload)

	// User is authenticated so let's create the access token (paseto)
	accessToken, err := rep.App.Paseto.CreateToken(user.Email, user.Nickname, user.Role)
	if err != nil {
		rep.errorJSON(w, errors.New("cannot create paseto token"), http.StatusUnauthorized)
		return
	}

	// Encode token into a secure cookie
	token.SetCookie(accessToken, w)

	// All set, redirect to home page with secure cookie in response header
	http.Redirect(w, r, "http://localhost/home", http.StatusSeeOther)

}
