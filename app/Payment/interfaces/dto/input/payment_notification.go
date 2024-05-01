package input

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type (
	PaymentNotificationInputDTO struct {
		ID               string              `json:"id"                validate:"required"`
		Action           string              `json:"action"`
		ApiVersion       string              `json:"api_version"`
		LiveMode         bool                `json:"live_mode"`
		UserID           string              `json:"user_id"`
		NotificationType string              `json:"type"`
		DateCreated      time.Time           `json:"date_created"`
		Data             PaymentDataInputDTO `json:"data"              validate:"required"`
	}

	PaymentDataInputDTO struct {
		ID string `json:"id"`
	}
)

func (p *PaymentNotificationInputDTO) ValidateDTO() error {
	return validator.New().Struct(p)
}
func (p *PaymentDataInputDTO) ValidateDTO() error {
	return validator.New().Struct(p)
}
