package token

import (
	"encoding/base64"
	"net/http"
)

// SetCookie encode token into a httponly secure cookie
func SetCookie(token string, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     CookieName,
		Value:    encodeCookie(token),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)
}

// Clear the cookie
func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   CookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

// Renew the cookie
func RenewCookie(token string, w http.ResponseWriter) {
	ClearSession(w)
	SetCookie(token, w)
}

// Encode the cookie value using base64.
func encodeCookie(s string) string {
	return base64.URLEncoding.EncodeToString([]byte(s))
}

// Decode the base64-encoded cookie value.
// Returns decoded cookie value (which holds encrypted token)
func DecodeCookie(cookie string) (paseto string, err error) {
	if value, err := base64.URLEncoding.DecodeString(cookie); err == nil {
		return string(value), nil
	}

	return "", err

}
