package api

import (
	"code-space-backend-api/api/middleware"
	"code-space-backend-api/app/Mail/usecases"
	"code-space-backend-api/app/config"
	"code-space-backend-api/common/cache"
	"code-space-backend-api/env"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(
	container *config.Container,
) Server {
	server := new(server)

	server.ginEngine = gin.Default()
	server.setupMiddlewares(container)
	server.setupHttpServer(env.Server.Port)

	server.setupRoutes(container)

	return server
}

func (s *server) setupMiddlewares(container *config.Container) {
	container.Metrics.Gin.Setup(s.ginEngine)

	if env.Tracer.IsTracingEnabled {
		s.ginEngine.Use(container.Tracing.Middleware)
	} else {
		s.ginEngine.Use(middleware.NewTraceMiddleware().Middleware())
		s.ginEngine.Use(middleware.NewRequestLoggerMiddleware(container.Log).Middleware())
	}

	s.ginEngine.Use(middleware.NewErrorMiddleware(container.Log).Middleware())
	s.cache = middleware.NewCacheMiddleware(cache.GetGlobalCacheManager())

	s.email = middleware.NewEmailMiddleware(
		usecases.NewSendEmailWithTemplate(
			usecases.NewSendEmail(
				container.EmailService,
			),
		),
	)

	s.token = middleware.NewTokenMiddleware()
}

func (s *server) setupHttpServer(port string) {
	tcpAddr := fmt.Sprintf(":%s", port)
	s.httpServer = &http.Server{
		Addr:    tcpAddr,
		Handler: s.ginEngine,
	}
}
