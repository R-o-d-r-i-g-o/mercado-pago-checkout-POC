package mapper

import (
	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/infra/database/models"
)

func ChapterModelsToOutputDTOs(models []models.Chapter) []output.ChapterOutputDTO {
	var outputDTOs []output.ChapterOutputDTO

	for _, model := range models {
		outputDTO := ChapterModelToOutputDTO(model)
		outputDTOs = append(outputDTOs, outputDTO)
	}

	return outputDTOs
}

func ChapterModelToOutputDTO(model models.Chapter) output.ChapterOutputDTO {
	return output.ChapterOutputDTO{
		Name:     model.Name,
		Hash:     model.Hash,
		Contents: mapContentsToOutputDTOs(model.Contents),
	}
}

func mapContentsToOutputDTOs(contents []models.Content) []output.ContentOutputDTO {
	var outputDTOs []output.ContentOutputDTO

	for _, content := range contents {
		outputDTOs = append(outputDTOs, ContentModelToOutputDTO(content))
	}

	return outputDTOs
}
