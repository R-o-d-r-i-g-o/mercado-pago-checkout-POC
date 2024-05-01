package email

import (
	"code-space-backend-api/app/Mail/entities"
	"time"

	"gopkg.in/gomail.v2"
)

func (sender *emailSender) SendEmail(email entities.Email) error {
	message := sender.composeEmailMessage(email)
	return sender.sendWithRetry(message)
}

func (sender *emailSender) composeEmailMessage(email entities.Email) *gomail.Message {
	message := gomail.NewMessage()

	if !email.IsValidContentType() {
		email = email.WithDefaultContentType()
	}

	message.SetBody(string(email.ContentType), email.Message)
	message.SetHeader("From", email.Sender)
	message.SetHeader("Subject", email.Subject)
	message.SetHeader("To", email.Receivers...)

	return message
}

func (sender *emailSender) sendWithRetry(message *gomail.Message) error {
	var err error
	for attempt := 1; attempt <= sender.maxAttempts; attempt++ {
		if err = sender.dialer.DialAndSend(message); err == nil {
			return nil
		}
		time.Sleep(sender.delay)
	}
	return err
}
