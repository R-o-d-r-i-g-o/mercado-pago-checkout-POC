package input

import "github.com/go-playground/validator/v10"

type ContentInputDTO struct {
	Hash string `json:"hash" validate:"required"`
}

func (input *ContentInputDTO) ValidateDTO() error {
	return validator.New().Struct(input)
}
