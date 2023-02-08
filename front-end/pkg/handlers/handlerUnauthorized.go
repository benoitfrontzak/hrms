package handlers

import (
	"frontend/pkg/render"
	"net/http"
)

// Unauthorized is the handler func rendering the unauthorized page
func (rep *Repository) Unauthorized(w http.ResponseWriter, r *http.Request) {
	// Empty data
	var data any
	// Render unauthorized template page
	render.RenderTemplate(w, "public.unauthorized.page.gohtml", data)
}
