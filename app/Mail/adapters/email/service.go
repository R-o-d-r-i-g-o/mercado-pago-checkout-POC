package email

import (
	"code-space-backend-api/app/Mail/entities"
	"time"

	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendEmail(email entities.Email) error
}

type emailSender struct {
	dialer *gomail.Dialer
	retryer
}

type retryer struct {
	maxAttempts int
	delay       time.Duration
}

func NewEmailService(host string, port int, username string, password string, maxAttempts int, delay time.Duration) EmailService {
	return &emailSender{
		dialer:  gomail.NewDialer(host, port, username, password),
		retryer: retryer{maxAttempts: maxAttempts, delay: delay},
	}
}
