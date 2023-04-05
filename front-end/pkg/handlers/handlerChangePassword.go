package handlers

import (
	"frontend/pkg/render"
	"log"
	"net/http"
)

func (rep *Repository) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Get my context from middleware request
	log.Println("At change password handler")
	myc := r.Context().Value(httpContext).(httpContextStruct)
	log.Println("Can't pass here ?")

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

	render.RenderTemplate(w, "admin.changePassword.page.gohtml", td)

}
