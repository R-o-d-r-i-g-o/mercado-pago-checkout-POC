package input

import (
	"code-space-backend-api/app/User/entities"

	"github.com/go-playground/validator/v10"
)

type (
	CreateUserInputDTO struct {
		Hash                  string `json:"hash"`
		Name                  string `json:"name" validate:"required,max=50"`
		Email                 string `json:"email" validate:"required,email"`
		Phone                 string `json:"phone" validate:"required,min=7"`
		Password              string `json:"password" validate:"required,min=6"`
		PolicyPrivacyAccepted bool   `json:"policy_privacy_accepted"`
	}
)

func (c *CreateUserInputDTO) ToDomain() entities.UserDomain {
	return entities.UserDomain{
		Hash:                  c.Hash,
		Name:                  c.Name,
		Email:                 c.Email,
		Phone:                 c.Phone,
		Password:              c.Password,
		PolicyPrivacyAccepted: c.PolicyPrivacyAccepted,
	}
}

func (c *CreateUserInputDTO) ValidateDTO() error {
	return validator.New().Struct(c)
}
