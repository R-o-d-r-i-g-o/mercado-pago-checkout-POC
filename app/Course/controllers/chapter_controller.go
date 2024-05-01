package controllers

import (
	"code-space-backend-api/api/response"
	"code-space-backend-api/app/Course/usecases"
	"code-space-backend-api/common/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChapterController interface {
	ListChapters(ctx *gin.Context)
	ListContentFromChapter(ctx *gin.Context)
}

type chapterController struct {
	listChapters           usecases.ListChapters
	listContentFromChapter usecases.ListContentFromChapter
}

func NewChapterController(listChapters usecases.ListChapters, listContentFromChapter usecases.ListContentFromChapter) ChapterController {
	return &chapterController{
		listChapters:           listChapters,
		listContentFromChapter: listContentFromChapter,
	}
}

func (c *chapterController) ListChapters(ctx *gin.Context) {
	courseHash := ctx.Param("course-hash")

	filter := pagination.NewPaginationFilterFromGinContext(ctx)

	chapterOutputDTO, err := c.listChapters.Execute(courseHash, filter)
	if err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	ctx.JSON(http.StatusFound, chapterOutputDTO)

}

func (c *chapterController) ListContentFromChapter(ctx *gin.Context) {
	chapterHash := ctx.Param("content-hash")

	chapterOuputDTO, err := c.listContentFromChapter.Execute(chapterHash)
	if err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	ctx.JSON(http.StatusFound, map[string]any{"module": chapterOuputDTO})
}
