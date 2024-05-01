package middleware

import (
	"context"

	"code-space-backend-api/common/constants/str"
	"code-space-backend-api/common/uuid"

	"github.com/gin-gonic/gin"
)

const traceHeaderName string = "x-trace"

type traceMiddleware struct{}

func NewTraceMiddleware() Middleware {
	return &traceMiddleware{}
}

func (t *traceMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		trace := c.Request.Header.Get(traceHeaderName)
		if trace == str.EMPTY_STRING {
			trace = uuid.New()
		}

		c.Set("trace", trace)

		ctx := context.WithValue(c.Request.Context(), traceHeaderName, trace)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
