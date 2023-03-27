package handlers

// RPCPayload is a struct which holds the information to be sent to the logger-service
type rpcPayload struct {
	Collection string
	Name       string
	Data       string
	CreatedAt  string
	CreatedBy  string
}

// jsonResponse is a struct which holds the response to be sent back
type jsonResponse struct {
	Error     bool   `json:"error"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
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
