package config

import (
	"code-space-backend-api/infra/database"

	"gorm.io/gorm"
)

type Infrastructure struct {
	DB *gorm.DB
}

func provideInfrastructure(c *Container) {
	c.Infrastructure = new(Infrastructure)
	if instance := database.GetInstance(); instance != nil {
		c.Infrastructure.DB = instance
	} else {
		c.Infrastructure.DB = database.Connect()
	}
}
