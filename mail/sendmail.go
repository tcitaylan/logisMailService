package mailer

import (
	"crypto/tls"
	"fmt"
	"time"

	gomail "gopkg.in/mail.v2"
)

func Smail() error {

	currentTime := time.Now()

	fmt.Printf("Sending eMail: %s", currentTime.Format("2006.01.02 15:04:05"))
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "logis8255@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "taylancivaoglu@gmail.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject: %s\n\n", currentTime.Format("2006.01.02 15:04:05"))

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "logis8255@gmail.com", "rwyy zvfs aeon sumc")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		//panic(err)
		return err
	}

	return nil
}
