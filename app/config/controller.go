package config

import (
	course "code-space-backend-api/app/Course/controllers"
	course_usecase "code-space-backend-api/app/Course/usecases"
	email "code-space-backend-api/app/Mail/controllers"
	email_usecase "code-space-backend-api/app/Mail/usecases"
	payment "code-space-backend-api/app/Payment/controllers"
	mercadopago "code-space-backend-api/app/Payment/gateway/mercado_pago"
	payment_usecase "code-space-backend-api/app/Payment/usecases"
	user "code-space-backend-api/app/User/controllers"
	user_usecase "code-space-backend-api/app/User/usecases"
	"code-space-backend-api/env"
	"net/http"
	"time"
)

type Controllers struct {
	User    user.UserController
	Email   email.EmailController
	Course  course.CourseController
	Chapter course.ChapterController
	Comment course.CommentController
	Payment payment.PaymentController
}

func provideControllers(c *Container) {
	c.Controllers = new(Controllers)

	provideUserConnect(c)
	provideEmailConnect(c)
	provideCourseConnect(c)
	provideChapterConnect(c)
	provideCommentConnect(c)
	providePaymentConnect(c)
}

func provideUserConnect(c *Container) {
	c.Controllers.User = user.NewUserController(
		user_usecase.NewCreateUser(c.Gateways.User),
		user_usecase.NewLoginUser(c.Gateways.User),
	)
}

func provideEmailConnect(c *Container) {

	c.Controllers.Email = email.NewEmailController(
		email_usecase.NewSendEmail(c.EmailService),
	)
}

func provideCourseConnect(c *Container) {
	c.Controllers.Course = course.NewCourseController(
		course_usecase.NewListCourses(c.Gateways.Course),
	)
}

func provideChapterConnect(c *Container) {
	c.Controllers.Chapter = course.NewChapterController(
		course_usecase.NewListChapters(c.Gateways.Chapter),
		course_usecase.NewListContentFromChapter(c.Gateways.Chapter),
	)
}

func provideCommentConnect(c *Container) {
	c.Controllers.Comment = course.NewCommentController(
		course_usecase.NewListComments(c.Gateways.Comment),
		course_usecase.NewGetCommentByHash(c.Gateways.Comment),
		course_usecase.NewCreateComment(c.Gateways.Comment),
	)
}

func providePaymentConnect(c *Container) {
	var service mercadopago.MercadoPagoService = mercadopago.NewMercadoPagoService(
		env.MercadoPago.BaseURL,
		env.MercadoPago.Token,
		http.DefaultClient,
	)

	const defaultRetryDelay = 2 * time.Minute
	const defaultRetryTimes = 5

	c.Controllers.Payment = payment.NewPaymentController(
		payment_usecase.NewCreatePaymentLink(c.Gateways.Purchase, service),
		payment_usecase.NewSearchPaymentInfo(service),
		payment_usecase.NewWaitForPaymentStatus(
			c.Gateways.Notification,
			c.Gateways.Purchase,
			c.Gateways.Payment,
			payment_usecase.NewSearchPaymentInfo(service),
			defaultRetryTimes,
			defaultRetryDelay,
		),
	)
}
