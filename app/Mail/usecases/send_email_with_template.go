package usecases

import (
	"code-space-backend-api/app/Mail/entities"
	"code-space-backend-api/app/Mail/interfaces/dto"
	"code-space-backend-api/common/constants/str"
	"os"
	"strings"
)

type SendEmailWithTemplate interface {
	Execute(emailDTO dto.LoginEmailTemplateInputDTO) error
}

type sendEmailWithTemplate struct {
	sendEmail SendEmail
}

func NewSendEmailWithTemplate(sendEmail SendEmail) SendEmailWithTemplate {
	return &sendEmailWithTemplate{
		sendEmail: sendEmail,
	}
}

func (usecase *sendEmailWithTemplate) Execute(emailDTO dto.LoginEmailTemplateInputDTO) error {
	if err := emailDTO.ValidateDTO(); err != nil {
		return invalidEmailMessageInput.WithArgs(err.Error())
	}

	email, err := usecase.createEmailMessage(emailDTO)
	if err != nil {
		return err
	}

	return usecase.sendEmail.Execute(dto.EmailInputDTO{
		ContentType: string(email.ContentType),
		Receivers:   email.Receivers,
		Message:     email.Message,
		Subject:     email.Subject,
		Sender:      email.Sender,
	})
}

func (usecase *sendEmailWithTemplate) createEmailMessage(emailDTO dto.LoginEmailTemplateInputDTO) (entities.Email, error) {
	var email entities.Email = emailDTO.ToDomain()

	message, err := usecase.loadEmailTemplate()
	if err != nil {
		return email, err
	}

	message = strings.ReplaceAll(message, "#USERNAME", emailDTO.Login)
	message = strings.ReplaceAll(message, "#PASSWORD", emailDTO.Password)

	return email.WithMessage(message), nil
}

func (usecase *sendEmailWithTemplate) loadEmailTemplate() (string, error) {
	root, err := os.Getwd()
	if err != nil {
		return str.EMPTY_STRING, err
	}

	htmlContent, err := os.ReadFile(root + "/app/Mail/templates/login.html")
	if err != nil {
		return str.EMPTY_STRING, err
	}

	return string(htmlContent), nil
}
