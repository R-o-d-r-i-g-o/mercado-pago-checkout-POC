package repository

import (
	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/app/Course/interfaces/mapper"
	"code-space-backend-api/common/pagination"
	"code-space-backend-api/infra/database/models"

	"gorm.io/gorm"
)

type CourseRepository interface {
	ListCourses(userHash string, filter pagination.PaginationFilter) ([]output.CourseOutputDTO, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) ListCourses(userHash string, filter pagination.PaginationFilter) ([]output.CourseOutputDTO, error) {
	var courses []models.Course

	offset, limit := filter.ToDatabaseFormat()

	err := r.db.Model(&models.Course{}).
		Joins("INNER JOIN user_courses uc ON courses.id = uc.course_id").
		Joins("INNER JOIN users u ON u.id = uc.user_id").
		Where("u.hash = ?", userHash).
		Offset(offset).
		Limit(limit).
		Find(&courses).
		Error

	if err != nil {
		return nil, err
	}

	return mapper.CourseModelsToOutputDTOs(courses), nil
}
