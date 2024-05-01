package usecases

import (
	"code-space-backend-api/app/Mail/adapters/email"
	"code-space-backend-api/app/Mail/interfaces/dto"
	"code-space-backend-api/common/errors"
)

const sendEmailContext string = "usecase/send-email"

var (
	invalidEmailMessageInput = errors.InvalidField.
					WithContext(sendEmailContext).
					WithName("invalid_field_send").
					WithTemplate("field send is either required or invalid: %v")

	failedToSendEmail = errors.Unknown.
				WithContext(sendEmailContext).
				WithName("failed_to_send_email").
				WithTemplate("error when sending email: %v")
)

type SendEmail interface {
	Execute(emailDTO dto.EmailInputDTO) error
}

type sendEmail struct {
	emailService email.EmailService
}

func NewSendEmail(emailService email.EmailService) SendEmail {
	return &sendEmail{
		emailService: emailService,
	}
}

func (usecase *sendEmail) Execute(emailDTO dto.EmailInputDTO) error {
	if err := emailDTO.ValidateDTO(); err != nil {
		return invalidEmailMessageInput.WithArgs(err.Error())
	}

	var message = emailDTO.ToDomain()

	if err := usecase.emailService.SendEmail(message); err != nil {
		return failedToSendEmail.WithArgs(err.Error())
	}

	return nil
}
