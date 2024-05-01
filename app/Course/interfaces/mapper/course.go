package mapper

import (
	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/infra/database/models"
)

func CourseModelsToOutputDTOs(models []models.Course) []output.CourseOutputDTO {
	var outputDTOs []output.CourseOutputDTO

	for _, model := range models {
		outputDTO := CourseModelToOutputDTO(model)
		outputDTOs = append(outputDTOs, outputDTO)
	}

	return outputDTOs
}

func CourseModelToOutputDTO(model models.Course) output.CourseOutputDTO {
	return output.CourseOutputDTO{
		Name:        model.Name,
		Hash:        model.Hash.String(),
		Description: model.Description,
		Modules:     mapModulesToOutputDTOs(model.Modules),
	}
}

func mapModulesToOutputDTOs(modules []models.Chapter) []output.ChapterOutputDTO {
	var outputDTOs []output.ChapterOutputDTO

	for _, module := range modules {
		outputDTOs = append(outputDTOs, ChapterModelToOutputDTO(module))
	}

	return outputDTOs
}
