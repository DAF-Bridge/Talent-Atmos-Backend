package repository

type MailRepository interface {
	SendInvitedMail(email string, subject string, Inviter string, context string) error
}
