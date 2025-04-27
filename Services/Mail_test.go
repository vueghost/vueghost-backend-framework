package Services

import "testing"

func TestNewMailSmtpClient(t *testing.T) {
	_ = NewMailSmtpClient()
}

func TestMailSmtpClient_Send(t *testing.T) {
	mailServer := MailSmtpClient{}
	message := MailMessage{
		From:     "no-reply@vueghost.com",
		FromName: "Test No Reply Email",
		To:       "mohamed_ym@yahoo.com",
		Subject:  "Test email subject",
		Content:  "Test email content",
	}
	if sent := mailServer.Send(message); !sent {
		t.Fail()
	}
}
