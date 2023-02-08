package handlers

import (
	"frontend/pkg/render"
	"net/http"
)

// Home is the handler func executed after middleware validation
// it holds a context sent by the middleware
func (rep *Repository) Mailer(w http.ResponseWriter, r *http.Request) {
	// Get my context from middleware request
	myc := r.Context().Value(httpContext).(httpContextStruct)

	// TODO: think how to share mailer page with other account...
	// So far anyone will have base.admin layout...
	if myc.Auth {
		// Render mailer page
		render.RenderTemplate(w, "admin.mailer.page.gohtml", myc.User)
	} else {
		var empty any
		// Render unauthorized page
		render.RenderTemplate(w, "unauthorized.page.gohtml", empty)
	}
}
