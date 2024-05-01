package usecases

import (
	"code-space-backend-api/app/Course/interfaces/dto/input"
	"code-space-backend-api/app/Course/interfaces/repository"
	"code-space-backend-api/common/constants/str"
	"code-space-backend-api/common/errors"
	"code-space-backend-api/common/token"
	"code-space-backend-api/common/uuid"
	"context"
)

const createCommentContext = "usecase/create-comment"

var (
	invalidCreateCommentInput = errors.InvalidField.
					WithContext(createCommentContext).
					WithName("invalid_field_send").
					WithTemplate("field send is either required or invalid: %v")

	failedToCreateComment = errors.Unknown.
				WithContext(createCommentContext).
				WithName("failed-to-create-comment").
				WithTemplate("could not create comment: %v")
)

type CreateComment interface {
	Execute(ctx context.Context, commentDTO input.CreateCommentInputDTO) error
}

type createComment struct {
	commentRepository repository.CommentRepository
}

func NewCreateComment(contentRepository repository.CommentRepository) CreateComment {
	return &createComment{
		commentRepository: contentRepository,
	}
}

func (usecase *createComment) Execute(ctx context.Context, commentDTO input.CreateCommentInputDTO) error {
	if err := commentDTO.ValidateDTO(); err != nil {
		return invalidCreateCommentInput.WithArgs(err.Error())
	}

	if commentDTO.Hash == str.EMPTY_STRING {
		commentDTO.Hash = uuid.New()
	}

	claims, err := token.ExtractTokenClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	commentDTO.UserHash = claims.CustomKeys["user_hash"].(string)

	if err := usecase.commentRepository.CreateComment(commentDTO); err != nil {
		return failedToCreateComment.WithArgs(err.Error())
	}

	return nil
}
