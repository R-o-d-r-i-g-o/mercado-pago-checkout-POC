package controller

import (
	"code-space-backend-api/api/response"
	"code-space-backend-api/app/Payment/interfaces/dto/input"
	"code-space-backend-api/app/Payment/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController interface {
	CreatePaymentLink(ctx *gin.Context)
	SearchPaymentByHash(ctx *gin.Context)
	ListenForPaymentNotifications(ctx *gin.Context)
}

type paymentController struct {
	createPaymentLink          usecases.CreatePaymentLink
	searchPaymentInfo          usecases.SearchPaymentInfo
	processPaymentNotification usecases.WaitForPaymentStatus
}

func NewPaymentController(createPaymentLink usecases.CreatePaymentLink, searchPaymentInfo usecases.SearchPaymentInfo, processPaymentNotification usecases.WaitForPaymentStatus) PaymentController {
	return &paymentController{
		createPaymentLink:          createPaymentLink,
		searchPaymentInfo:          searchPaymentInfo,
		processPaymentNotification: processPaymentNotification,
	}
}

func (c *paymentController) CreatePaymentLink(ctx *gin.Context) {
	var paymentInfo input.PaymentLinkInputDTO

	if err := ctx.ShouldBindJSON(&paymentInfo); err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	paymentLink, err := c.createPaymentLink.Execute(ctx, paymentInfo)
	if err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusCreated, response.Data(paymentLink))
}

func (c *paymentController) SearchPaymentByHash(ctx *gin.Context) {
	var paymentHash string = ctx.Param("paymentHash")

	paymentInfo, err := c.searchPaymentInfo.Execute(paymentHash)
	if err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusFound, paymentInfo)
}

func (c *paymentController) ListenForPaymentNotifications(ctx *gin.Context) {
	var notification input.PaymentNotificationInputDTO

	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	if err := c.processPaymentNotification.Execute(notification); err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
