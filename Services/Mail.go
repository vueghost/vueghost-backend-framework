package Services

import "gopkg.in/gomail.v2"

const MailSmtpHost string = "email-smtp.us-west-2.8987888.com"
const MailSmtpPort int = 587
const MailSmtpUsername string = ""
const MailSmtpPassword string = "+"

type (
	//MailSmtpClient Permits email to be sent using Mail, Sendmail, or SMTP.
	MailSmtpClient struct {
	}
	//MailMessage Mail message struct.
	MailMessage struct {
		From     string
		FromName string
		To       string
		Subject  string
		Content  string
	}
)

func NewMailSmtpClient() *MailSmtpClient {
	return &MailSmtpClient{}
}

func (c MailSmtpClient) Send(message MailMessage) (sentSuccessfully bool) {
	smtpMail := gomail.NewMessage()
	smtpMail.SetHeaders(map[string][]string{
		"From":    {smtpMail.FormatAddress(message.From, message.FromName)},
		"To":      {message.To},
		"Subject": {message.Subject},
	})

	smtpMail.SetBody("text/html", message.Content)
	smtpMailDialer := gomail.NewDialer(MailSmtpHost, MailSmtpPort, MailSmtpUsername, MailSmtpPassword)
	if err := smtpMailDialer.DialAndSend(smtpMail); err != nil {
		return false
	}

	return true
}
