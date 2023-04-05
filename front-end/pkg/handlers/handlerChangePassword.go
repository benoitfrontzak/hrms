package handlers

import (
	"frontend/pkg/render"
	"net/http"
)

func (rep *Repository) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Get my context from middleware request
	myc := r.Context().Value(httpContext).(httpContextStruct)

	// If unauthorized request or unexpected request method
	if !myc.Auth {
		var empty any
		render.RenderTemplate(w, "public.unauthorized.page.gohtml", empty)
		return
	}

	// store template data
	td := TemplateData{
		User: myc.User,
	}
	switch myc.User.Role {
	case 2:
		render.RenderTemplate(w, "admin.changePassword.page.gohtml", td)
	case 4:
		render.RenderTemplate(w, "user.changePassword.page.gohtml", td)
	}

}
