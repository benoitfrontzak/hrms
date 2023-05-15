package handlers

// jsonResponse is a struct which holds the response to be sent back
type jsonResponse struct {
	Error     bool   `json:"error"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
	CreatedAt string `json:"createdAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
}

// is the struct which holds a new entry of CT to be added
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

// listID holds the list of id to be deleted
type listID struct {
	List          []string `json:"list,omitempty"`
	Amount        int      `json:"amount,string,omitempty"`
	Reason        string   `json:"reason,omitempty"`
	ConnectedUser string   `json:"connectedUser,omitempty"`
	UserID        int      `json:"userID,string,omitempty"`
}

// User holds the connected user information (id, email)
type User struct {
	ID    int    `json:"id,string,omitempty"`
	Email string `json:"email,omitempty"`
}
