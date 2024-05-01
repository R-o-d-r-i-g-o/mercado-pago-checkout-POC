package usecases

import (
	"code-space-backend-api/app/Course/interfaces/dto/input"
	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/app/Course/interfaces/repository"
	"code-space-backend-api/common/errors"
	"code-space-backend-api/common/pagination"
)

const listCommentsContext string = "usecase/list-comments"

var (
	invalidContentInput = errors.InvalidField.
				WithContext(listCommentsContext).
				WithName("invalid_field_send").
				WithTemplate("field send is either required or invalid: %v")

	failedToListCommentsWithGivenContentHash = errors.Unknown.
							WithContext(listCommentsContext).
							WithName("failed_comments_given_content_hash").
							WithTemplate("could not list comments given content hash: %v")
)

type ListComments interface {
	Execute(contentDTO input.ContentInputDTO, filter pagination.PaginationFilter) (pagination.Page[output.CommentOutputDTO], error)
}

type listComments struct {
	commentRepository repository.CommentRepository
}

func NewListComments(contentRepository repository.CommentRepository) ListComments {
	return &listComments{
		commentRepository: contentRepository,
	}
}

func (usecase *listComments) Execute(contentDTO input.ContentInputDTO, filter pagination.PaginationFilter) (pagination.Page[output.CommentOutputDTO], error) {
	var response pagination.Page[output.CommentOutputDTO]

	if err := contentDTO.ValidateDTO(); err != nil {
		return response, invalidContentInput.WithArgs(err.Error())
	}

	comments, err := usecase.commentRepository.ListCommentsGivenContentHash(contentDTO.Hash, filter)
	if err != nil {
		return response, failedToListCommentsWithGivenContentHash.WithArgs(err.Error())
	}

	return response.With(comments), nil
}
