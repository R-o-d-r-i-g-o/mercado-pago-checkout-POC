package input

import "github.com/go-playground/validator/v10"

type CreateCommentInputDTO struct {
	Text        string `json:"text" validate:"required"`
	Hash        string `json:"hash"`
	UserHash    string `json:"-"`
	ParentHash  string `json:"parent_hash"`
	ContentHash string `json:"content_hash" validate:"required"`
}

func (input *CreateCommentInputDTO) ValidateDTO() error {
	return validator.New().Struct(input)
}
