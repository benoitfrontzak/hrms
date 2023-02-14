package handlers

import (
	"frontend/pkg/render"
	"log"
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
	if !myc.Auth {
		var empty any
		// Render unauthorized page
		render.RenderTemplate(w, "unauthorized.page.gohtml", empty)
		return
	}

	// Check request method (GET|POST)
	switch r.Method {

	case "GET":
		employeeGET(w, myc, r.URL.Path)

	case "POST":
		employeePOST(w, myc, r.URL.Path)

	default:
		log.Println("Sorry, only GET & POST method are supported.")
		return
	}

}

func employeeGET(w http.ResponseWriter, myc httpContextStruct, url string) {
	// Get expected action from URL (/employee/delete/)
	sURL := strings.Split(url, "/")
	action := sURL[2]

	switch action {
	case "add":
		showEmployeeAddForm(w, myc)
	case "update":
		// showEmployeeEditForm(w, r)
	case "delete":
		// deleteEmployeeByID(w, r)
	default:
		showEmployeeHome(w, myc)
	}
}
func showEmployeeHome(w http.ResponseWriter, myc httpContextStruct) {
	// Role: 1: superAdmin; 2 : admin; 3: Manager; 4: User
	td := TemplateData{
		User: myc.User,
	}
	switch myc.User.Role {
	// Admin employee home page
	case 2:
		render.RenderTemplate(w, "admin.employee.page.gohtml", td)

	// User employee home page
	case 3:
		// Fetch employee profile (TODO)

		render.RenderTemplate(w, "user.employee.page.gohtml", td)
	}
}
func showEmployeeAddForm(w http.ResponseWriter, myc httpContextStruct) {

	// Role: 1: superAdmin; 2 : admin; 3: Manager; 4: User
	td := TemplateData{
		User: myc.User,
	}
	switch myc.User.Role {
	// Admin add employee page
	case 2:
		render.RenderTemplate(w, "admin.employee.add.page.gohtml", td)

	}
}
func employeePOST(w http.ResponseWriter, myc httpContextStruct, url string) {
	sURL := strings.Split(url, "/")
	action := sURL[2]

	switch action {
	// case "add":
	// 	showEmployeeAddForm(w, myc)
	// case "update":
	// 	showEmployeeEditForm(w, r)
	// case "delete":
	// 	deleteEmployeeByID(w, r)
	// default:
	// 	showEmployeeHome(w, myc)
	}
}

// func fetchAllUsers() (resp *http.Response, err error) {
// 	resp, err = http.Get("http://localhost:8081/allUsers")
// 	if err != nil {
// 		log.Println("error while query http.Get: ", err)
// 		return nil, err
// 	}
// 	return
// }

// Fetch all employees (TO put inside handler)
// usr, _ := fetchAllUsers()
// var p jsonResponse
// err := json.NewDecoder(usr.Body).Decode(&p)
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusBadRequest)
// 	return
// }
