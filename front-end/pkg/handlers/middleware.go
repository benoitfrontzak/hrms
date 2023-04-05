package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	unauthorized = "http://localhost/unauthorized"
	authorize    = "http://localhost:8081/api/v1/authentication/verifyCookie"
)

// Middleware checks if secure cookie exists
// if the cookie is valid, send it to authentication-service (resp = ResponseCookieVerify)
// if resp.Status = http.StatusAccepted (200) set httpContext to request and forward it
// if resp.Status !=  http.StatusAccepted redirect to unauthorized page
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get secure cookie from request
		cookie, err := r.Cookie(CookieName)
		if err != nil {
			http.Redirect(w, r, unauthorized, http.StatusSeeOther)
			return
		}

		// verifyCookiePayload is the struct which holds the encoded cookie
		// to be sent to authentication-service for validation
		verifyCookiePayload := VerifyCookiePayload{
			Cookie: *cookie,
		}

		// Get authentication-service response
		resp, err := verifyCookiePayload.verifyCookie()
		if err != nil || resp.Status != 200 {
			log.Printf("failed to verify the cookie: %s", err)
			http.Redirect(w, r, unauthorized, http.StatusSeeOther)
			return
		}

		// Set httpContextStruct
		r = r.WithContext(context.WithValue(
			r.Context(),
			httpContext,
			httpContextStruct{
				Auth: true,
				User: resp.TokenPayload,
			},
		))

		// Authorized request
		next.ServeHTTP(w, r)

	})
}

// Verify cookie's authorization and validity
// send post request to authentication-service
func (p *VerifyCookiePayload) verifyCookie() (r *ResponseCookieVerify, err error) {
	// Set form parameters
	params := url.Values{}
	params.Add("cookie", p.Cookie.Value)

	// Get authentication-service response
	resp, err := http.PostForm(authorize, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal response body to ResponseCookieVerify
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
