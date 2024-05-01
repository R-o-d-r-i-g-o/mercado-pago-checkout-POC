package config

import (
	"code-space-backend-api/app/Mail/adapters/email"
	"code-space-backend-api/common/logger"
	"code-space-backend-api/env"
	"fmt"
	"log"
	"time"
)

type Container struct {
	ServiceName        string
	ServiceEnvironment string
	EmailService       email.EmailService
	Log                logger.LogEngine
	Tracing            *Tracing
	Metrics            *Metrics
	Controllers        *Controllers
	Gateways           *Gateways
	Infrastructure     *Infrastructure
}

func NewContainer() *Container {
	c := new(Container)
	provide(c)
	return c
}

func provide(c *Container) {
	provideServiceInfo(c)
	provideLog(c)
	provideEmailService(c)
	provideTracing(c)
	provideMetrics(c)

	provideInfrastructure(c)
	provideGateways(c)
	provideControllers(c)
}

func provideServiceInfo(c *Container) {
	c.ServiceName = fmt.Sprintf("%s-%s",
		env.GeneralConfig.ServiceName,
		env.GeneralConfig.ServiceEnvironment,
	)
	c.ServiceEnvironment = env.GeneralConfig.ServiceEnvironment
}

func provideEmailService(c *Container) {
	const defaultRetryTimes int = 5
	const defaultRetryDelay time.Duration = 2 * time.Minute

	c.EmailService = email.NewEmailService(
		env.EmailService.Host,
		env.EmailService.Port,
		env.EmailService.Username,
		env.EmailService.Password,
		defaultRetryTimes,
		defaultRetryDelay,
	)
}

func provideLog(c *Container) {
	var err error

	c.Log, err = logger.New(env.Logger.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
}
