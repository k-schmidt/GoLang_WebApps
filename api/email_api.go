package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

const (
	EmailAddress  = ""
	EmailPassword = ""
)

type email_struct struct {
	EmailRecipients []string
	EmailSubject    string
	EmailMessage    string
}

func send(from string, password string, t email_struct) {
	msg := "From: " + from + "\n" +
		"To: " + strings.Join(t.EmailRecipients, ",") + "\n" +
		"Subject: " + t.EmailSubject + "\n" +
		t.EmailMessage

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, t.EmailRecipients, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Println("sent")
}

func EmailApi(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t email_struct
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(t.EmailMessage)
	send(EmailAddress, EmailPassword, t)
}

// main ...
func main() {
	http.HandleFunc("/", EmailApi)
	http.ListenAndServe(":8080", nil)
}
