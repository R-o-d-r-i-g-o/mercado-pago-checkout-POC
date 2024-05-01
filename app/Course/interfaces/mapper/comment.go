package mapper

import (
	"code-space-backend-api/app/Course/interfaces/dto/input"
	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/infra/database/models"
)

func CreateCommentInputDtoToModel(input input.CreateCommentInputDTO) models.Comment {
	return models.Comment{
		Text: input.Text,
		Hash: input.Hash,
	}
}

func ModelCommentsToOutputCommentsDTOs(model []models.Comment) []output.CommentOutputDTO {
	var commentsDTO []output.CommentOutputDTO

	for _, comment := range model {
		commentsDTO = append(commentsDTO, ModelCommentToOutputCommentDTO(comment))
	}

	return commentsDTO
}

func ModelCommentToOutputCommentDTO(model models.Comment) output.CommentOutputDTO {
	commentDTO := output.CommentOutputDTO{
		Hash:        model.Hash,
		Content:     model.Text,
		PublishedAt: model.CreatedAt,
		Author: output.AuthorOutputDTO{
			Name: model.Author.Name,
		},
		Children: []output.CommentOutputDTO{},
	}

	for _, child := range model.Children {
		commentDTO.Children = append(commentDTO.Children, ModelCommentToOutputCommentDTO(child))
	}

	return commentDTO
}
