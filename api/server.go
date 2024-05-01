package api

import (
	"code-space-backend-api/api/middleware"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Server interface {
	Run()
	Shutdown(ctx context.Context)
}

type server struct {
	ginEngine  *gin.Engine
	httpServer HttpServer
	cache      middleware.Middleware
	email      middleware.Middleware
	token      middleware.Middleware
}

func (s *server) Run() {
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func (s *server) Shutdown(ctx context.Context) {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
