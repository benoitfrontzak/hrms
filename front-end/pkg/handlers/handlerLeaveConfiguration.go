package handlers

import (
	"frontend/pkg/render"
	"net/http"
	"strings"
)

// LeaveConfiguration is the handler func executed after middleware validation
// it holds a context sent by the middleware
func (rep *Repository) LeaveConfiguration(w http.ResponseWriter, r *http.Request) {
	// Check URL
	if !strings.HasPrefix(r.URL.Path, "/configuration/leave/") {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Get my context from middleware request
	myc := r.Context().Value(httpContext).(httpContextStruct)
	// If unauthorized request or unexpected request method
	if !myc.Auth || r.Method != "GET" {
		var empty any
		render.RenderTemplate(w, "unauthorized.page.gohtml", empty)
		return
	}

	// store template data
	td := TemplateData{
		User: myc.User,
	}

	// Role: 1=superAdmin; 2=admin; 3=Manager; 4=User
	// only admin is allowed to managed employee Config Tables
	if myc.User.Role == 2 {
		render.RenderTemplate(w, "admin.leaveCT.page.gohtml", td)
	} else {
		var empty any
		render.RenderTemplate(w, "unauthorized.page.gohtml", empty)
	}

}
