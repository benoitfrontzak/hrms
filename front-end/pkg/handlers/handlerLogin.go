package handlers

import (
	"frontend/pkg/render"
	"net/http"
)

func (rep *Repository) Login(w http.ResponseWriter, r *http.Request) {
	// Build business logic
	var data interface{}
	// Render template page
	render.RenderTemplate(w, "public.login.page.gohtml", data)
}
