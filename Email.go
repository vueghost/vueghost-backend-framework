package System

import (
	"Framework/Services"
)

//Email type struct.
type Email struct {
	mailSmtpClient Services.MailSmtpClient
	mailMessage    Services.MailMessage
	view           View
}

//Context Email context.
func (e *Email) Context() {
	e.mailMessage = Services.MailMessage{
		From:     "no-replay@vueghost.com",
		FromName: "vueghost",
	}
}

//SetTo set message must send to email address.
func (e *Email) SetTo(emailTo string) {
	e.mailMessage.To = emailTo
}

//SetSubject Set email message subject.
func (e *Email) SetSubject(subject string) {
	e.mailMessage.Subject = subject
}

//SetMessage Set email message.
func (e *Email) SetMessage(message string) {
	e.mailMessage.Content = message
}

func (e *Email) ViewRender(viewName string, data interface{}) string {
	return e.view.ViewRender(viewName, data)
}

//Send Send email message using mail service.
func (e *Email) Send() bool {
	mail := Services.NewMailSmtpClient()
	return mail.Send(e.mailMessage)
}
