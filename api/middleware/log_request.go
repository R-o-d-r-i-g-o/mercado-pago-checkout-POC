package middleware

import (
	"bytes"
	"io"

	"code-space-backend-api/common/logger"

	"github.com/gin-gonic/gin"
)

type requestLoggerMiddleware struct {
	log logger.LogEngine
}

func NewRequestLoggerMiddleware(log logger.LogEngine) Middleware {
	return &requestLoggerMiddleware{log: log}
}

func (r *requestLoggerMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := r.log.WithContext(c.Request.Context())

		log.Info("incoming request",
			logger.String("clientIp", c.ClientIP()),
			logger.String("url", c.Request.URL.String()),
			logger.Any("headers", c.Request.Header),
			logger.BytesString("body", r.getBody(c)),
		)

		c.Next()
	}
}

func (r *requestLoggerMiddleware) getBody(c *gin.Context) []byte {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	body, _ := io.ReadAll(tee)
	c.Request.Body = io.NopCloser(&buf)

	return body
}
