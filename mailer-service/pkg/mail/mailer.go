/*
	 _____ _  _ _   _ _  _ ___  ___ ___  ___  ___  ___ _____
	|_   _| || | | | | \| |   \| __| _ \/ __|/ _ \| __|_   _|
	  | | | __ | |_| | .` | |) | _||   /\__ \ (_) | _|  | |
	  |_| |_||_|\___/|_|\_|___/|___|_|_\|___/\___/|_|   |_|

1.	Package name: 							mail
2.	Date of creation:						2023-02-02
3.	Name of creator of module:				Benoit Frontzak
4.	History of modification:				N/A (v1)
5.	Summary of what the module does:		Mail format email into sendgrid personalization template format, and send it to sendgrid's API
6.	Functions in that module:				SendToAPI; CreateListEmails; DynamicTemplateEmail
7.	Variables accessed by the module:		subject, content string, recipients, carbonCopys []*mail.Email

Important: more details about the personalization email template can be found:

	https://github.com/sendgrid/sendgrid-go/blob/main/examples/helpers/mail/example.go
*/
package mail

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Sendgrid API configuration
const (
	keyAPI   = "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"
	endpoint = "/v3/mail/send"
	host     = "https://api.sendgrid.com"

	address  = "thundersoft.my.hrms@gmail.com" // must the a valid sendgrid account address
	name     = "HRMS"
	replyToN = "noreply"
	replyToA = "noreply@thundersoft.com"
)

// SendToAPI send an email to sendgrid API
// it takes a personalized sengrid's template as parameter
func SendToAPI(body []byte) error {
	// Set request parameters
	request := sendgrid.GetRequest(keyAPI, endpoint, host)
	request.Method = "POST"
	request.Body = body

	// Send request to sendgrid API
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println("error while sending email through API", err)
		return err
	}

	log.Println("email successfully sent though API, response status code: ", response.StatusCode)

	return nil
}

// CreateListEmails make a list of email sendgrid's format
func CreateListEmails(e []string) []*mail.Email {
	//  My emails address
	var myEA = make([]*mail.Email, 0, len(e))

	for _, email := range e {
		new := mail.NewEmail("", email)
		myEA = append(myEA, new)
	}

	return myEA
}

// TODO: make func for attachment
// a := mail.NewAttachment()
// a.SetContent("TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4gQ3JhcyBwdW12")
// a.SetType("application/pdf")
// a.SetFilename("balance_001.pdf")
// a.SetDisposition("attachment")
// m.AddAttachment(a)

// DynamicTemplateEmail creates an email template for sendgrid
// it returns a request body to be used by sendgrid API
func DynamicTemplateEmail(subject, content string, recipients, carbonCopys []*mail.Email) []byte {
	// Create new email
	m := mail.NewV3Mail()

	// Set email set from
	e := mail.NewEmail(name, address)
	m.SetFrom(e)

	// Set email subject
	m.Subject = subject

	// Create new email personalization
	p := mail.NewPersonalization()

	// Add all recipients to email
	tos := recipients
	p.AddTos(tos...)

	// Add all carbonCopys to email
	ccs := carbonCopys
	p.AddCCs(ccs...)

	// Add html content to email
	c := mail.NewContent("text/html", content)
	m.AddContent(c)

	// Set reply to
	replyToEmail := mail.NewEmail(replyToN, replyToA)
	m.SetReplyTo(replyToEmail)

	// Add personalization to email
	m.AddPersonalizations(p)

	return mail.GetRequestBody(m)
}
