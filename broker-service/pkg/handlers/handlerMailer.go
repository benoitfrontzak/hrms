package handlers

import (
	"broker/pkg/producer"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type payloadMail struct {
	Recipients string `json:"email_to"`
	Ccs        string `json:"email_cc"`
	Subject    string `json:"email_subject"`
	Message    string `json:"email_message"`
}

// My queue is the queue name declare by amqp-service
var myQ = "mailer"

// SendEmail receive a post request from the front-end
// it publishes payloadMail to rabbitMQ mailer queue
func (rep *Repository) SendEmail(w http.ResponseWriter, r *http.Request) {
	// Invoke ParseForm before read form values
	r.ParseForm()

	// Extract payload from request
	p := payloadMail{
		Recipients: r.FormValue("email_to"),
		Ccs:        r.FormValue("email_cc"),
		Subject:    r.FormValue("email_subject"),
		Message:    r.FormValue("email_message"),
	}

	// Publishes the payload to Rabbit mailer queue
	// The queue is declare by amqp-service
	err := rep.pushToQueue(p, myQ)
	if err != nil {
		log.Println("error while publishing the queue mailer: ", err)
	}

	// Log email publishing to logger service
	uEmail := r.FormValue("user_email")
	uNick := r.FormValue("user_nickname")

	rpcPayload := RPCPayload{
		Collection: "notification",
		Name:       "email successfully published to " + myQ + " queue",
		Data:       fmt.Sprintf("%s (%s) sent an email to %s with subject %s", uNick, uEmail, r.FormValue("email_to"), r.FormValue("email_subject")),
	}
	rep.LogItemViaRPC(w, rpcPayload)

	log.Printf("email payload successfully published to %s queue", myQ)

	http.Redirect(w, r, "http://localhost/home", http.StatusSeeOther)
}

// pushToQueue publishes a message into RabbitMQ queue
func (r *Repository) pushToQueue(p payloadMail, queue string) error {
	// Create new producer type
	producer := producer.NewProducer(r.App.Rabbit)

	// Format payload ([]byte)
	j, err := json.MarshalIndent(&p, "", "\t")
	if err != nil {
		return err
	}
	// Publishing the message
	err = producer.PushMail(queue, j)
	if err != nil {
		return err
	}

	return nil
}
