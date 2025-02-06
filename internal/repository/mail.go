package repository

type MailRepository interface {
	SendMail(email string, subject string, Inviter string, context string) error
}
