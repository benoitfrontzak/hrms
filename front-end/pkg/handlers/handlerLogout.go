package handlers

import (
	"log"
	"net/http"
)

// Logout is the handler func which log out the user and clear the cookie
func (rep *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	clearCookie(w)
	http.Redirect(w, r, "http://localhost", http.StatusSeeOther)
}

// Clear the cookie
func clearCookie(w http.ResponseWriter) {
	log.Printf("%s %s %s \n", "the cookie", CookieName, "have been deleted successfully")
	cookie := &http.Cookie{
		Name:   CookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
