package controllers

import (
	"code-space-backend-api/api/response"
	"code-space-backend-api/app/Mail/interfaces/dto"
	"code-space-backend-api/app/Mail/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailController interface {
	SendEmail(ctx *gin.Context)
}

type emailController struct {
	sendEmail usecases.SendEmail
}

func NewEmailController(sendEmail usecases.SendEmail) EmailController {
	return &emailController{
		sendEmail: sendEmail,
	}
}

func (c *emailController) SendEmail(ctx *gin.Context) {
	var email dto.EmailInputDTO

	if err := ctx.ShouldBindJSON(&email); err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	if err := c.sendEmail.Execute(email); err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
