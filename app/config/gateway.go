package config

import (
	course "code-space-backend-api/app/Course/interfaces/repository"
	payment "code-space-backend-api/app/Payment/interfaces/repository"
	user "code-space-backend-api/app/User/interfaces/repository"
)

type Gateways struct {
	User         user.UserRepository
	Course       course.CourseRepository
	Chapter      course.ChapterRepository
	Comment      course.CommentRepository
	Payment      payment.PaymentRepository
	Purchase     payment.PurchaseRepository
	Notification payment.NotificationRepository
}

func provideGateways(c *Container) {
	c.Gateways = new(Gateways)
	provideUserRepository(c)
	provideCourseRepository(c)
	provideChapterRepository(c)
	provideCommentRepository(c)
	providePaymentRepository(c)
	providePurchaseRepository(c)
	provideNotificationRepository(c)
}

func provideUserRepository(c *Container) {
	c.Gateways.User = user.NewUserRepository(c.Infrastructure.DB)
}

func provideCourseRepository(c *Container) {
	c.Gateways.Course = course.NewCourseRepository(c.Infrastructure.DB)
}

func provideChapterRepository(c *Container) {
	c.Gateways.Chapter = course.NewChapterRepository(c.Infrastructure.DB)
}

func provideCommentRepository(c *Container) {
	c.Gateways.Comment = course.NewCommentRepository(c.Infrastructure.DB)
}

func providePurchaseRepository(c *Container) {
	c.Gateways.Purchase = payment.NewPurchaseRepository(c.Infrastructure.DB)
}

func providePaymentRepository(c *Container) {
	c.Gateways.Payment = payment.NewPaymentRepository(c.Infrastructure.DB)
}

func provideNotificationRepository(c *Container) {
	c.Gateways.Notification = payment.NewNotificationRepository(c.Infrastructure.DB)
}
