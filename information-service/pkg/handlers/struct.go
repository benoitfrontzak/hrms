package handlers

// RPCPayload is a struct which holds the information to be sent to the logger-service
type RPCPayload struct {
	Collection string
	Name       string
	Data       string
}

// AuthPayload is the embedded type (in RequestPayload) that describes an authentication request
type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// jsonResponse is a struct which holds the response to be sent back
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
