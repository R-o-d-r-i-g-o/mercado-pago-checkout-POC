package database

import "code-space-backend-api/infra/database/models"

func MigrateModels() {
	instance.AutoMigrate(&models.ContentType{})
	instance.AutoMigrate(&models.VideoContent{})
	instance.AutoMigrate(&models.User{})
	instance.AutoMigrate(&models.Course{})
	instance.AutoMigrate(&models.UserCourse{})
	instance.AutoMigrate(&models.Chapter{})
	instance.AutoMigrate(&models.Content{})
	instance.AutoMigrate(&models.Comment{})
	instance.AutoMigrate(&models.UserProgress{})
	instance.AutoMigrate(&models.PaymentStatus{})
	instance.AutoMigrate(&models.PaymentType{})
	instance.AutoMigrate(&models.PaymentMethod{})
	instance.AutoMigrate(&models.Product{})
	instance.AutoMigrate(&models.Payment{})
	instance.AutoMigrate(&models.Notification{})
	instance.AutoMigrate(&models.Purchase{})
	instance.AutoMigrate(&models.PurchaseProduct{})
}
