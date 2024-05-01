package api

import (
	course "code-space-backend-api/app/Course/controllers"
	email "code-space-backend-api/app/Mail/controllers"
	payment "code-space-backend-api/app/Payment/controllers"
	user "code-space-backend-api/app/User/controllers"
	"code-space-backend-api/app/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) setupRoutes(container *config.Container) {
	if s.ginEngine == nil {
		log.Fatal("engine is nil")
	}

	groupV1 := s.ginEngine.Group("/v1")
	groupApi := groupV1.Group("/api")

	s.setupHealthCheck(s.ginEngine)
	s.setupUserConnectRouter(groupApi, container.Controllers.User)
	s.setuptMailGunConnectRouter(groupApi, container.Controllers.Email)
	s.setupCourseConnectRouter(groupApi, container.Controllers.Course)
	s.setupChapterConnectRouter(groupApi, container.Controllers.Chapter)
	s.setupCommentConnectRouter(groupApi, container.Controllers.Comment)
	s.setupMercadoPagoConnectRouter(groupApi, container.Controllers.Payment)

}

func (s *server) setupUserConnectRouter(engine *gin.RouterGroup, c user.UserController) {
	usersGroup := engine.Group("/users")
	usersGroup.POST("/create", s.email.Middleware(), c.Create)
	usersGroup.POST("/login", c.Login)
}

func (s *server) setupCourseConnectRouter(engine *gin.RouterGroup, c course.CourseController) {
	courseGroup := engine.Group("/courses", s.token.Middleware())
	courseGroup.GET("listcourses", c.ListCourses)
}

func (s *server) setupChapterConnectRouter(engine *gin.RouterGroup, c course.ChapterController) {
	chapterGroup := engine.Group("/chapters", s.token.Middleware())
	chapterGroup.GET("/listcontents/:content-hash", c.ListContentFromChapter)
	chapterGroup.GET("/listchapters/:course-hash", c.ListChapters)
}

func (s *server) setupCommentConnectRouter(engine *gin.RouterGroup, c course.CommentController) {
	commentGroup := engine.Group("/comment", s.token.Middleware())
	commentGroup.GET("/getcomment/:comment-hash", c.GetCommentWithChildrenByHash)
	commentGroup.GET("/listcomments/:content-hash", c.ListContentComments)
	commentGroup.POST("/create", c.CreateComment)
}

func (s *server) setupMercadoPagoConnectRouter(engine *gin.RouterGroup, c payment.PaymentController) {
	mercadoPagoGroup := engine.Group("/mercadopago")
	mercadoPagoGroup.POST("/create-payment-link", c.CreatePaymentLink)
	mercadoPagoGroup.GET("/search-payment/:paymentHash", c.SearchPaymentByHash)
	mercadoPagoGroup.POST("/payment-notification-webhook", c.ListenForPaymentNotifications)
}

func (s *server) setuptMailGunConnectRouter(engine *gin.RouterGroup, c email.EmailController) {
	mailGunGroup := engine.Group("/emailservice")
	mailGunGroup.POST("/sendmail", c.SendEmail)
}

func (s *server) setupHealthCheck(engine *gin.Engine) {
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "it's running!",
		})
	})
}
