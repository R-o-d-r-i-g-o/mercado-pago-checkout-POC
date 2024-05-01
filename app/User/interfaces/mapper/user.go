package mapper

import (
	"code-space-backend-api/app/User/interfaces/dto/input"
	"code-space-backend-api/app/User/interfaces/dto/output"
	"code-space-backend-api/infra/database/models"
)

func UserModelToOutputDTO(model models.User) output.UserOutputDTO {
	return output.UserOutputDTO{
		Hash:     model.Hash,
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	}
}

func UserInputDtoToModel(inputDTO input.CreateUserInputDTO) models.User {
	return models.User{
		Hash:                  inputDTO.Hash,
		Name:                  inputDTO.Name,
		Email:                 inputDTO.Email,
		Phone:                 inputDTO.Phone,
		Password:              inputDTO.Password,
		PolicyPrivacyAccepted: inputDTO.PolicyPrivacyAccepted,
	}
}
