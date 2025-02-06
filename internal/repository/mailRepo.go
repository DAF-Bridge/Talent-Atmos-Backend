package repository

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"sync"
)

type InviteMailRepository struct {
	mailserver *gomail.Dialer
}

var (
	tmpl     *template.Template
	loadOnce sync.Once
)

func loadTemplate() {
	loadOnce.Do(func() {
		var err error
		tmpl, err = template.ParseFiles("Invite_email_template.html")
		if err != nil {
			log.Fatalf("Error loading template: %v", err)
		}
	})
}

func NewInviteMailRepository(mailserver *gomail.Dialer) *InviteMailRepository {
	return &InviteMailRepository{mailserver: mailserver}
}

func (i *InviteMailRepository) SendMail(email string, subject string, Inviter string, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", i.mailserver.Username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	context, err := makeHtmlInviteBody(Inviter, token)
	if err != nil {
		return err
	}
	m.SetBody("text/html", context)
	return i.mailserver.DialAndSend(m)
}

func makeHtmlInviteBody(Inviter string, token string) (string, error) {
	loadTemplate()

	data := struct {
		Inviter string
		Token   string
		URL     string
	}{
		Inviter: Inviter,
		Token:   token,
		URL:     "https://www.google.com/",
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil

}
