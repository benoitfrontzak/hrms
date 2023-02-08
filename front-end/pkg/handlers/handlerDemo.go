package handlers

import (
	"encoding/json"
	"frontend/pkg/render"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Firstname string    `json:"firstname,omitempty"`
	Lastname  string    `json:"lastname,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Password  string    `json:"-"`
	Active    int       `json:"active"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// TerminatedAt string    `json:"terminated_at"`
}
type jsonResponse struct {
	Data []User
}
type templateData struct {
	Nickname string
	Data     []User
}

// Home is the handler func executed after middleware validation
// it holds a context sent by the middleware
func (rep *Repository) Demo(w http.ResponseWriter, r *http.Request) {
	// Get my context from middleware request
	myc := r.Context().Value(httpContext).(httpContextStruct)

	// Fetch all users as demo
	usr, _ := fetchAllUsers()
	var p jsonResponse
	err := json.NewDecoder(usr.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	td := templateData{
		Nickname: myc.User.Nickname,
		Data:     p.Data,
	}
	// TODO: think how to share mailer page with other account...
	// So far anyone will have base.admin layout...
	if myc.Auth {
		// Render mailer page
		render.RenderTemplate(w, "admin.demo.page.gohtml", td)
	} else {
		var empty any
		// Render unauthorized page
		render.RenderTemplate(w, "unauthorized.page.gohtml", empty)
	}
}

func fetchAllUsers() (resp *http.Response, err error) {
	resp, err = http.Get("http://localhost:8081/allUsers")
	if err != nil {
		log.Println("error while query http.Get: ", err)
		return nil, err
	}
	return
}
