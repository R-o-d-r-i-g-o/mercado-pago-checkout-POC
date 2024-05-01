package response

import (
	"code-space-backend-api/common/errors"

	"github.com/go-playground/validator/v10"
)

var ErrInvalidFields = errors.Validation.
	WithName("invalid_fields").
	WithMessage("invalid fields found in body")

var ErrBadFormat = errors.Validation.
	WithName("bad_format")

func ComposeError(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		validationError := ErrInvalidFields

		for _, errField := range validationErrors {
			invalidFieldError := errors.InvalidField.
				WithParam(errField.Namespace()).
				WithMessage("field is %v", errField.Tag())

			validationError = validationError.Add(invalidFieldError)
		}

		return validationError
	}

	return ErrBadFormat.WithMessage("something wrong on format: %w", err)
}
