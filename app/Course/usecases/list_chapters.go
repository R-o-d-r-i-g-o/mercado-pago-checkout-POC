package usecases

import (
	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/app/Course/interfaces/repository"
	"code-space-backend-api/common/errors"
	"code-space-backend-api/common/pagination"
)

const listChapterContext = "usecase/list-all-modules"

var (
	couldNotFindModulesWithGivenCourseHash = errors.NotFound.
		WithContext(listChapterContext).
		WithName("could_not_find_modules_with_given_course_hash").
		WithTemplate("failed to get modules associated to the course: %v")
)

type ListChapters interface {
	Execute(courseHash string, filter pagination.PaginationFilter) (pagination.Page[output.ChapterOutputDTO], error)
}

type listChapters struct {
	repository repository.ChapterRepository
}

func NewListChapters(chapterRepository repository.ChapterRepository) ListChapters {
	return &listChapters{
		repository: chapterRepository,
	}
}

func (usecase *listChapters) Execute(courseHash string, filter pagination.PaginationFilter) (pagination.Page[output.ChapterOutputDTO], error) {
	var responsePaginated pagination.Page[output.ChapterOutputDTO]

	modules, err := usecase.repository.ListChaptersByCourseHash(courseHash, filter)
	if err != nil {
		return responsePaginated, couldNotFindModulesWithGivenCourseHash.WithArgs(err.Error())
	}

	return responsePaginated.With(modules), nil
}
