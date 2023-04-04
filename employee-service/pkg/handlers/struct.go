package handlers

// AuthPayload is the embedded type (in RequestPayload) that describes an authentication request
type authPayload struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Fullname   string `json:"fullname"`
	Nickname   string `json:"nickname"`
	Active     int    `json:"active"`
	Role       int    `json:"role"`
	EmployeeID int    `json:"employeeID"`
}

// jsonResponse is a struct which holds the response to be sent back
type jsonResponse struct {
	Error     bool   `json:"error"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
}

// newEmployee is a struct which holds the information of the created employee
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

// employeeListID is the struct which holds the list of employee id to be deleted
type employeeListID struct {
	List          []string `json:"list,omitempty"`
	ConnectedUser string   `json:"connectedUser,omitempty"`
}

//  is the struct which holds a new entry of CT to be added
type configTablePayload struct {
	Table         string `json:"name,omitempty"`
	Entry         string `json:"value,omitempty"`
	ConnectedUser string `json:"email,omitempty"`
	RowID         int    `json:"rowid,string,omitempty"`
}
type ctEntriesPayload struct {
	List          []string `json:"list,omitempty"`
	Table         string   `json:"name,omitempty"`
	ConnectedUser string   `json:"email,omitempty"`
}
