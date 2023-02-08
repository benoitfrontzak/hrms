package handlers

import (
	"authentication/pkg/token"
	"encoding/json"
	"log"
	"net/http"
)

// threshold is the time on which we renew the cookie/token
// var threshold = 10 * time.Minute

// VerifyCookie verifies the cookie authenticity and validity (expired date)
// if successfull, it returns http.StatusOK and user information
// if failed, it returns http.StatusUnauthorized
func (rep *Repository) VerifyCookie(w http.ResponseWriter, r *http.Request) {
	// Response to front-end
	var verifyCookieResponse = VerifyCookieResponse{}

	// Extract paseto token from request
	p, err := rep.extractToken(r)
	if err != nil {
		log.Printf("error while extracting token: %s", err)
		return
	}

	// Let's assume first that the token still valid
	verifyCookieResponse = VerifyCookieResponse{
		Status:       http.StatusOK,
		Message:      "authorized connexion",
		TokenPayload: p,
	}

	// Check if token expired and returns duration
	// Update response accordingly
	// d, err := p.Valid()
	err = p.Valid()
	if err != nil {
		// log.Printf("Token expired for user %s, autorenew it", p.Nickname)
		// verifyCookieResponse = VerifyCookieResponse{
		// 	Status:  http.StatusUnauthorized,
		// 	Message: "unauthorized connexion, token is expired",
		// }
		verifyCookieResponse = rep.autorenew(p, w)
	}

	// Check if token is about to expires
	// if d < threshold {
	// 	verifyCookieResponse = rep.autorenew(p, w)
	// }

	json.NewEncoder(w).Encode(&verifyCookieResponse)

}

// extractToken extracts the encoded cookie and returns the decrypted token
func (rep *Repository) extractToken(r *http.Request) (*token.TokenData, error) {
	// Get Secure Cookie from post request
	// Invoke ParseForm before read form values
	r.ParseForm()
	sc := r.FormValue("cookie")

	// Decode encoded cookie (holds paseto Token)
	t, err := token.DecodeCookie(sc)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return nil, err
	}
	// Decode paseto into token.Payload
	p, err := rep.App.Paseto.DecryptToken(t)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return nil, err
	}

	return p, nil
}

// autorenew creates a new paseto token and renew the cookie with it
func (rep *Repository) autorenew(p *token.TokenData, w http.ResponseWriter) VerifyCookieResponse {
	// Recreate token string
	ts, _ := rep.App.Paseto.CreateToken(p.Email, p.Nickname, p.Role)

	// Renew cookie (clear & create)
	token.RenewCookie(ts, w)

	// Decode new token
	nt, err := rep.App.Paseto.DecryptToken(ts)
	if err != nil {
		log.Println("error while decrypting the new token: ", err)
	}

	verifyCookieResponse := VerifyCookieResponse{
		Status:       http.StatusOK,
		Message:      "authorized connexion",
		TokenPayload: nt,
	}

	return verifyCookieResponse
}
