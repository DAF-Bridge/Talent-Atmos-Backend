package initializers

import (
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
)

var DialerMail *gomail.Dialer

func SetupMail() {
	//SMTP_PASSWORD
	//SMTP_MAIL
	//SMTP_HOST
	//SMTP_PORT
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal(err.Error())
	}
	smtpMail := os.Getenv("SMTP_MAIL")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	//check if the smtp variables are set
	if smtpHost == "" || smtpMail == "" || smtpPassword == "" {
		log.Fatal("SMTP variables are not set")
	}

	DialerMail = gomail.NewDialer(smtpHost, smtpPort, smtpMail, smtpPassword)

}
