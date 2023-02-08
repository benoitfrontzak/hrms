package handlers

// jsonResponse is a struct which holds the response to be sent back
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// payloadMail is the stuct which holds the payload sent from broker-service
type payloadMail struct {
	Recipients string `json:"email_to"`
	Ccs        string `json:"email_cc"`
	Subject    string `json:"email_subject"`
	Message    string `json:"email_message"`
}
