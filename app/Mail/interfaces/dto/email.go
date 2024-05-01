package dto

import (
	"code-space-backend-api/app/Mail/entities"

	"github.com/go-playground/validator/v10"
)

type EmailInputDTO struct {
	Receivers   []string `json:"receivers" validate:"required"`
	ContentType string   `json:"content_type" validate:"required"`
	Message     string   `json:"message" validate:"required"`
	Subject     string   `json:"subject" validate:"required"`
	Sender      string   `json:"sender" validate:"required"`
}

func (e *EmailInputDTO) ValidateDTO() error {
	return validator.New().Struct(e)
}

func (e *EmailInputDTO) ToDomain() entities.Email {
	return entities.Email{
		ContentType: entities.EmailMessageType(e.ContentType),
		Receivers:   e.Receivers,
		Message:     e.Message,
		Subject:     e.Subject,
		Sender:      e.Sender,
	}
}
