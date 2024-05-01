package controllers

import (
	"code-space-backend-api/api/response"
	"code-space-backend-api/app/Course/interfaces/dto/input"
	"code-space-backend-api/app/Course/usecases"
	"code-space-backend-api/common/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	ListContentComments(ctx *gin.Context)
	GetCommentWithChildrenByHash(ctx *gin.Context)
}

type commentController struct {
	listComments     usecases.ListComments
	getCommentByHash usecases.GetCommentByHash
	createComment    usecases.CreateComment
}

func NewCommentController(listComments usecases.ListComments, getCommentByHash usecases.GetCommentByHash,
	createComment usecases.CreateComment,
) CommentController {
	return &commentController{
		listComments:     listComments,
		getCommentByHash: getCommentByHash,
		createComment:    createComment,
	}
}

func (c *commentController) CreateComment(ctx *gin.Context) {
	var comment input.CreateCommentInputDTO

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	if err := c.createComment.Execute(ctx, comment); err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	ctx.AbortWithStatus(http.StatusCreated)
}

func (c *commentController) ListContentComments(ctx *gin.Context) {
	hash := ctx.Param("content-hash")

	contentDTO := input.ContentInputDTO{
		Hash: hash,
	}

	filter := pagination.NewPaginationFilterFromGinContext(ctx)

	comments, err := c.listComments.Execute(contentDTO, filter)
	if err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusFound, map[string]any{"comments": comments})
}

func (c *commentController) GetCommentWithChildrenByHash(ctx *gin.Context) {
	hash := ctx.Param("comment-hash")

	contentDTO := input.ContentInputDTO{
		Hash: hash,
	}

	comment, err := c.getCommentByHash.Execute(contentDTO)
	if err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusFound, comment)
}
