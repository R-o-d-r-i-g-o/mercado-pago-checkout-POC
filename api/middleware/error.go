package middleware

import (
	"net/http"

	"code-space-backend-api/common/errors"
	"code-space-backend-api/common/logger"

	"code-space-backend-api/api/response"

	"github.com/gin-gonic/gin"
)

type errorMiddleware struct {
	log logger.LogEngine
}

func NewErrorMiddleware(log logger.LogEngine) Middleware {
	return &errorMiddleware{log: log}
}

func (e *errorMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		log := e.log.WithContext(c.Request.Context())

		if !e.shouldHandleErrors(c) {
			return
		}

		e.logErrors(c, log)

		lastError := c.Errors.Last()
		err, ok := lastError.Err.(errors.Error)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		e.handleErrorCode(c, err)
	}
}

func (e *errorMiddleware) shouldHandleErrors(c *gin.Context) bool {
	return len(c.Errors) > 0
}

func (e *errorMiddleware) logErrors(c *gin.Context, log logger.LogEngine) {
	for _, ginErr := range c.Errors {
		log.Error(ginErr.Err)
	}
}

func (e *errorMiddleware) handleErrorCode(c *gin.Context, err errors.Error) {
	statusCode := errors.HttpStatusCode(err)

	if statusCode == http.StatusInternalServerError {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatusJSON(statusCode, response.Error(err))
}
