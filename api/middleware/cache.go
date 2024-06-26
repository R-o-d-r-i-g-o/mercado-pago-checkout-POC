package middleware

import (
	"bytes"
	"code-space-backend-api/common/cache"
	"code-space-backend-api/common/constants/str"
	"code-space-backend-api/common/token"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	defaultCacheTimeout time.Duration = time.Minute * 20
	cacheHeaderKey      string        = "X-Cache-Control"
	noCache             string        = "no-cache"
)

type cacheMiddleware struct {
	cacheManager cache.CacheManager
}

type cacheWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func NewCacheMiddleware(cacheManager cache.CacheManager) Middleware {
	return &cacheMiddleware{cacheManager: cacheManager}
}

func (c *cacheMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if c.shouldHandleRequest(ctx) {
			ctx.Next()
			return
		}

		cacheKey, err := c.createCacheKeyFromRequest(ctx)
		if err != nil {
			ctx.Next()
			return
		}

		if cache, exists := c.cacheManager.Get(cacheKey); exists {
			ctx.AbortWithStatusJSON(http.StatusFound, c.parseCachedValue(cache))
			return
		}

		cacheWriter := newCacheWriter(ctx.Writer)
		ctx.Writer = cacheWriter

		ctx.Next()
		c.cacheResponse(cacheKey, cacheWriter.body.Bytes())
	}
}

func newCacheWriter(w gin.ResponseWriter) *cacheWriter {
	return &cacheWriter{ResponseWriter: w, body: bytes.NewBuffer(nil)}
}

func (w *cacheWriter) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

func (c *cacheMiddleware) shouldHandleRequest(ctx *gin.Context) bool {
	if ctx.Request.Method != http.MethodGet {
		return false
	}

	value := ctx.GetHeader(cacheHeaderKey)
	return strings.ToLower(value) == noCache
}

func (c *cacheMiddleware) parseCachedValue(value interface{}) interface{} {
	var dataMap map[string]interface{}
	fmt.Println(string(value.([]byte)))
	if err := json.Unmarshal(value.([]byte), &dataMap); err == nil {
		return dataMap
	}

	var dataArray []map[string]interface{}
	if err := json.Unmarshal(value.([]byte), &dataArray); err == nil {
		return dataArray
	}

	return nil
}

func (c *cacheMiddleware) cacheResponse(key string, body []byte) {
	c.cacheManager.Set(key, body, defaultCacheTimeout)
}

func (c *cacheMiddleware) createCacheKeyFromRequest(ctx *gin.Context) (string, error) {
	userHash, err := c.getUserHashFromTokenClaims(ctx)
	if err != nil {
		return str.EMPTY_STRING, err
	}

	return fmt.Sprintf("%s%s", userHash, ctx.Request.URL), nil
}

func (c *cacheMiddleware) getUserHashFromTokenClaims(ctx *gin.Context) (string, error) {
	claims, err := token.ExtractTokenClaimsFromContext(ctx)
	if err != nil {
		return str.EMPTY_STRING, err
	}

	hash, ok := claims.CustomKeys["user_hash"].(string)
	if !ok {
		return hash, errors.New("no hash founded")
	}

	return hash, nil
}
