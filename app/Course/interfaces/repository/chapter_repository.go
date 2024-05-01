package repository

import (
	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/app/Course/interfaces/mapper"
	"code-space-backend-api/common/pagination"
	"code-space-backend-api/infra/database/models"

	"gorm.io/gorm"
)

type chapterRepository struct {
	db *gorm.DB
}

type ChapterRepository interface {
	ListContentByChapterHash(moduleHash string) (output.ChapterOutputDTO, error)
	ListChaptersByCourseHash(courseHash string, filter pagination.PaginationFilter) ([]output.ChapterOutputDTO, error)
}

func NewChapterRepository(db *gorm.DB) ChapterRepository {
	return &chapterRepository{db: db}
}

func (r *chapterRepository) ListContentByChapterHash(moduleHash string) (output.ChapterOutputDTO, error) {
	var module models.Chapter

	err := r.db.Model(models.Chapter{}).
		Preload("Contents").
		Preload("Contents.VideoContent").
		Where("hash = ?", moduleHash).
		First(&module).
		Error

	if err != nil {
		return output.ChapterOutputDTO{}, err
	}

	return mapper.ChapterModelToOutputDTO(module), nil
}

func (r *chapterRepository) ListChaptersByCourseHash(courseHash string, filter pagination.PaginationFilter) ([]output.ChapterOutputDTO, error) {
	var modules []models.Chapter

	offset, limit := filter.ToDatabaseFormat()
	err := r.db.Model(&models.Chapter{}).
		Joins("INNER JOIN courses ON courses.id = chapters.course_id").
		Where("courses.hash = ?", courseHash).
		Offset(offset).
		Find(&modules).
		Limit(limit).
		Error

	if err != nil {
		return nil, err
	}

	return mapper.ChapterModelsToOutputDTOs(modules), nil
}
