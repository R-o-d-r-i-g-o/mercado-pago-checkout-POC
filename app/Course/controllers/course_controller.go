package controllers

import (
	"code-space-backend-api/api/response"
	"code-space-backend-api/app/Course/usecases"
	"code-space-backend-api/common/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseController interface {
	ListCourses(ctx *gin.Context)
}

type courseController struct {
	listCourses usecases.ListCourses
}

func NewCourseController(listCourses usecases.ListCourses) CourseController {
	return &courseController{
		listCourses: listCourses,
	}
}

func (c *courseController) ListCourses(ctx *gin.Context) {
	filter := pagination.NewPaginationFilterFromGinContext(ctx)

	outputDTO, err := c.listCourses.Execute(ctx, filter)
	if err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	ctx.JSON(http.StatusFound, outputDTO)
}
