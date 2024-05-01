package usecases

import (
	"code-space-backend-api/app/Course/interfaces/dto/input"
	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/app/Course/interfaces/repository"
	"code-space-backend-api/common/errors"
)

const getCommentByHashContext = "usecase/list-comments"

var (
	failedToSearchCommentWithGivenCommentHash = errors.Unknown.
		WithContext(getCommentByHashContext).
		WithName("failed_comments_given_content_hash").
		WithTemplate("could not list comments given content hash: %v")
)

type GetCommentByHash interface {
	Execute(contentDTO input.ContentInputDTO) (output.CommentOutputDTO, error)
}

type getCommentByHash struct {
	commentRepository repository.CommentRepository
}

func NewGetCommentByHash(commentRepository repository.CommentRepository) GetCommentByHash {
	return &getCommentByHash{
		commentRepository: commentRepository,
	}
}

func (usecase *getCommentByHash) Execute(contentDTO input.ContentInputDTO) (output.CommentOutputDTO, error) {
	comment, err := usecase.commentRepository.FindCommentWithChildren(contentDTO.Hash)
	if err != nil {
		return comment, failedToSearchCommentWithGivenCommentHash.WithArgs(err.Error())
	}

	return comment, nil
}
