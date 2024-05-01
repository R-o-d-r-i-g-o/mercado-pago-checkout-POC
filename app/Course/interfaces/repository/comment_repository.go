package repository

import (
	"code-space-backend-api/app/Course/interfaces/dto/input"
	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/app/Course/interfaces/mapper"
	"code-space-backend-api/common/pagination"
	"code-space-backend-api/infra/database/models"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

type CommentRepository interface {
	ListCommentsGivenContentHash(contentHash string, filter pagination.PaginationFilter) ([]output.CommentOutputDTO, error)
	FindCommentWithChildren(commentHash string) (output.CommentOutputDTO, error)
	CreateComment(commentDTO input.CreateCommentInputDTO) error
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) FindCommentWithChildren(commentHash string) (output.CommentOutputDTO, error) {
	var comment models.Comment

	err := r.db.Model(models.Comment{}).
		Preload("Author").
		Preload("Children.Author").
		Where("hash = ?", commentHash).
		First(&comment).
		Error

	if err != nil {
		return output.CommentOutputDTO{}, err
	}

	return mapper.ModelCommentToOutputCommentDTO(comment), nil
}

func (r *commentRepository) ListCommentsGivenContentHash(contentHash string, filter pagination.PaginationFilter) ([]output.CommentOutputDTO, error) {
	var comments []models.Comment

	offset, limit := filter.ToDatabaseFormat()
	err := r.db.Model(models.Comment{}).
		Joins("INNER JOIN contents c ON c.id = comments.content_id").Preload("Author").
		Where("comments.parent_id IS NULL").
		Where("c.hash = ?", contentHash).
		Order("comments.ID").
		Offset(offset).
		Limit(limit).
		Find(&comments).
		Error

	if err != nil {
		return nil, err
	}

	return mapper.ModelCommentsToOutputCommentsDTOs(comments), nil
}

func (r *commentRepository) CreateComment(commentDTO input.CreateCommentInputDTO) error {
	authorID, err := r.getAuthorID(commentDTO.UserHash)
	if err != nil {
		return err
	}

	parentID, err := r.getParentID(commentDTO.ParentHash)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	contentID, err := r.getContentID(commentDTO.ContentHash)
	if err != nil {
		return err
	}

	commentModel := mapper.CreateCommentInputDtoToModel(commentDTO)
	commentModel.AuthorID = authorID
	commentModel.ContentID = contentID

	if parentID != 0 {
		commentModel.ParentID = &parentID
	}

	return r.db.
		Model(models.Comment{}).
		Create(&commentModel).
		Error
}

func (r *commentRepository) getAuthorID(userHash string) (uint, error) {
	var authorID uint

	return authorID, r.db.
		Model(models.User{}).Select("id").
		Where("hash = ?", userHash).
		First(&authorID).
		Error
}

func (r *commentRepository) getParentID(parentHash string) (uint, error) {
	var parentID uint

	err := r.db.
		Model(models.Comment{}).Select("id").
		Where("hash = ?", parentHash).
		First(&parentID).
		Error

	if err == gorm.ErrRecordNotFound {
		return parentID, nil
	}

	return parentID, err
}

func (r *commentRepository) getContentID(contentHash string) (uint, error) {
	var contentID uint

	return contentID, r.db.
		Model(models.Content{}).Select("id").
		Where("hash = ?", contentHash).
		First(&contentID).
		Error
}
