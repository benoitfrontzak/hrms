package handlers

import (
	"frontend/pkg/render"
	"net/http"
	"strings"
)

// Employee is the handler func executed after middleware validation
// it holds a context sent by the middleware
func (rep *Repository) Employee(w http.ResponseWriter, r *http.Request) {
	// Check URL
	if !strings.HasPrefix(r.URL.Path, "/employee/") {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Get my context from middleware request
	myc := r.Context().Value(httpContext).(httpContextStruct)

	// If unauthorized request or unexpected request method
	if !myc.Auth || r.Method != "GET" {
		var empty any
		render.RenderTemplate(w, "public.unauthorized.page.gohtml", empty)
		return
	}

	// store template data
	td := TemplateData{
		User: myc.User,
	}

	// check action requested
	// action := strings.TrimPrefix(r.URL.Path, "/employee/")
	action := strings.Split(r.URL.Path, "/")[2]

	// Role: 1=superAdmin; 2=admin; 3=Manager; 4=User
	// only admin is allowed to managed employee
	if myc.User.Role == 2 {
		switch action {
		case "update":
			render.RenderTemplate(w, "admin.employee.update.page.gohtml", td)
		default:
			render.RenderTemplate(w, "admin.employee.page.gohtml", td)
		}
	} else {
		var empty any
		render.RenderTemplate(w, "public.unauthorized.page.gohtml", empty)
	}
}
