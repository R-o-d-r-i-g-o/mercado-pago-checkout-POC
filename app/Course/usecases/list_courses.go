package usecases

import (
	"context"

	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/app/Course/interfaces/repository"
	"code-space-backend-api/common/errors"
	"code-space-backend-api/common/pagination"
	"code-space-backend-api/common/token"
)

const listCoursesContext = "usecase/list-all-courses"

var (
	failedToListAllCourses = errors.Unknown.
		WithContext(listCoursesContext).
		WithName("failed_to_list_all_courses").
		WithTemplate("could not list all the courses: %v")
)

type ListCourses interface {
	Execute(ctx context.Context, filter pagination.PaginationFilter) (pagination.Page[output.CourseOutputDTO], error)
}

type listCourses struct {
	repository repository.CourseRepository
}

func NewListCourses(courseRepository repository.CourseRepository) ListCourses {
	return &listCourses{
		repository: courseRepository,
	}
}

func (usecase *listCourses) Execute(ctx context.Context, filter pagination.PaginationFilter) (response pagination.Page[output.CourseOutputDTO], err error) {
	claims, err := token.ExtractTokenClaimsFromContext(ctx)
	if err != nil {
		return response, err
	}

	var userHash = claims.CustomKeys["user_hash"].(string)

	courses, err := usecase.repository.ListCourses(userHash, filter)
	if err != nil {
		return response, failedToListAllCourses.WithArgs(err.Error())
	}

	response = response.With(courses)

	return
}
