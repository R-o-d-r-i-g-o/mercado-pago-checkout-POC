package config

import (
	"code-space-backend-api/common/tracing"

	"code-space-backend-api/env"

	"github.com/gin-gonic/gin"
)

type Tracing struct {
	Middleware gin.HandlerFunc
	Provider   tracing.Provider
}

func provideTracing(c *Container) {
	c.Tracing = new(Tracing)

	if !env.Tracer.IsTracingEnabled {
		c.Tracing.Provider = tracing.NewNopProvider()
		return
	}

	c.Tracing.Provider = tracing.SetupJaegerProvider(c.ServiceName, c.ServiceEnvironment)

	c.Tracing.Middleware = tracing.GinMiddleware()
}
