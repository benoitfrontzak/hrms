package handlers

import (
	"mailer/pkg/mail"
	"net/http"
	"strings"
)

// Mailer is the handler func receiving an email payload from the broker-service
// it sends email
func (rep *Repository) Mailer(w http.ResponseWriter, r *http.Request) {
	// Extract payload from request
	var rp payloadMail

	err := rep.readJSON(w, r, &rp)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// Convert recipients and carbon copys from payload (string) to []string
	// we assume that the front-end validated the appropriate format:
	// each email have to be separated by ;
	slctos := strings.Split(rp.Recipients, ";")
	slcccs := strings.Split(rp.Ccs, ";")

	// Convert list of emails to sendgrid emails list
	recipients := mail.CreateListEmails(slctos)
	carbonCopys := mail.CreateListEmails(slcccs)

	// Format request payload to sendgrid body's template
	myB := mail.DynamicTemplateEmail(rp.Subject, rp.Message, recipients, carbonCopys)

	// Send email to sendgrig API
	err = mail.SendToAPI(myB)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	rep.writeJSON(w, http.StatusAccepted, "")
}
