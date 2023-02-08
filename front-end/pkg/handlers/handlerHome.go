package handlers

import (
	"frontend/pkg/render"
	"net/http"
)

// Home is the handler func executed after middleware validation
// it holds a context sent by the middleware
func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// Get my context from middleware request
	myc := r.Context().Value(httpContext).(httpContextStruct)

	// Redirect to home page according to granted access (user role)
	myc.authorization(w)

}

// authorization redirect to the home page of the granted user acccess
func (r *httpContextStruct) authorization(w http.ResponseWriter) {
	authorized := r.Auth
	user := r.User

	// Authorized request
	if authorized {
		switch user.Role {
		case 1:
			// Render superAdminHome page
			render.RenderTemplate(w, "superAdmin.home.page.gohtml", user)

		case 2:
			// Render adminHome page
			render.RenderTemplate(w, "admin.home.page.gohtml", user)

		case 3:
			// Render managerHome page
			render.RenderTemplate(w, "manager.home.page.gohtml", user)

		case 4:
			// Render userHome page
			render.RenderTemplate(w, "user.home.page.gohtml", user)

		default:
			var empty any
			// Render unauthorized page
			render.RenderTemplate(w, "unauthorized.page.gohtml", empty)
		}

	} else {
		var empty any
		// Render unauthorized page
		render.RenderTemplate(w, "unauthorized.page.gohtml", empty)
	}

}
