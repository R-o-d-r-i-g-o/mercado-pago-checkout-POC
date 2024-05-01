package middleware

import (
	"code-space-backend-api/common/constants/str"
	"code-space-backend-api/common/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const claimsHeaderName string = "claims"

type tokenMiddleware struct{}

func NewTokenMiddleware() Middleware {
	return &tokenMiddleware{}
}

func (t *tokenMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := t.extractBearerToken(c.Request)
		if tokenString == str.EMPTY_STRING {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims, err := token.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set(claimsHeaderName, claims)

		c.Next()
	}
}

func (t *tokenMiddleware) extractBearerToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == str.EMPTY_STRING {
		return str.EMPTY_STRING
	}

	return strings.Split(authHeader, str.EMPTY_SPACE)[1]
}
