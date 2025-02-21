package repository

import (
	"bytes"
	"html/template"

	"gopkg.in/gomail.v2"
)

type InviteMailRepository struct {
	mailserver            *gomail.Dialer
	tmpl                  *template.Template
	baseCallbackInviteURL string
}

func NewInviteMailRepository(mailserver *gomail.Dialer, tmpl *template.Template, baseCallbackInviteURL string) MailRepository {
	return &InviteMailRepository{
		mailserver:            mailserver,
		tmpl:                  tmpl,
		baseCallbackInviteURL: baseCallbackInviteURL}
}

func (i *InviteMailRepository) SendInvitedMail(config InviteMailConfig) error {
	m := gomail.NewMessage()
	m.SetHeader("From", i.mailserver.Username)
	m.SetHeader("To", config.ToEmail)
	m.SetHeader("Subject", config.Subject)
	context, err := i.makeHtmlInviteBody(config.Body)
	if err != nil {
		return err
	}
	m.SetBody("text/html", context)
	return i.mailserver.DialAndSend(m)
}

func (i *InviteMailRepository) makeHtmlInviteBody(Body InviteMailBody) (string, error) {

	dataInTmpl := struct {
		User    string
		Inviter string
		URL     string
		ORG     string
	}{
		User:    Body.InvitedName,
		Inviter: Body.InviterName,
		URL:     i.baseCallbackInviteURL + Body.Token,
		ORG:     Body.OrganizationName,
	}

	var tpl bytes.Buffer
	if err := i.tmpl.Execute(&tpl, dataInTmpl); err != nil {
		return "", err
	}

	return tpl.String(), nil

}
