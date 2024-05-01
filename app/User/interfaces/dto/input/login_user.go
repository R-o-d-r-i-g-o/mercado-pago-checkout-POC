package input

import (
	"code-space-backend-api/app/User/entities"

	"github.com/go-playground/validator/v10"
)

type (
	LoginUserInputDTO struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}
)

func (l *LoginUserInputDTO) ToDomain() entities.UserDomain {
	return entities.UserDomain{
		Email:    l.Email,
		Password: l.Password,
	}
}

func (l *LoginUserInputDTO) ValidateDTO() error {
	return validator.New().Struct(l)
}
