package tracing

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func GinMiddleware() gin.HandlerFunc {
	return otelgin.Middleware(ServiceName())
}
