package mapper

import (
	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/infra/database/models"
)

func ContentModelToOutputDTO(model models.Content) output.ContentOutputDTO {
	return output.ContentOutputDTO{
		Hash:         model.Hash,
		TypeID:       uint(model.TypeID),
		VideoContent: VideoModelToOutputDTO(model.VideoContent),
	}
}

func VideoModelToOutputDTO(model models.VideoContent) output.VideoOutputDTO {
	return output.VideoOutputDTO{
		URL:         model.Url,
		Name:        model.Name,
		Description: model.Description,
	}
}
