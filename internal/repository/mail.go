package repository

type InviteMailConfig struct {
	ToEmail string
	Subject string
	Body    InviteMailBody
}

type InviteMailBody struct {
	InviterName      string
	InvitedName      string
	OrganizationName string
	Token            string
}

type MailRepository interface {
	SendInvitedMail(InviteMailConfig) error
}
