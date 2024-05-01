package token

import (
	"code-space-backend-api/common/constants/str"
	"code-space-backend-api/common/errors"
	"code-space-backend-api/env"
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const tokenContext = "common/token"

var (
	invalidToken = errors.BusinessRule.
			WithContext(tokenContext).
			WithName("invalid_token")

	invalidSignature = errors.Unathorizated.
				WithContext(tokenContext).
				WithName("unexpected_signing_method")

	failedToExtractTokenClaims = errors.Unknown.
					WithContext(tokenContext).
					WithName("failed_to_extract_token_claims")
)

type CustomClaims struct {
	jwt.StandardClaims
	CustomKeys map[string]any `json:"custom_claims,omitempty"`
}

func CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(env.GeneralConfig.SecretKey))
	if err != nil {
		return str.EMPTY_STRING, err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidSignature
		}
		return []byte(env.GeneralConfig.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, invalidToken
}

func ExtractTokenClaimsFromContext(ctx context.Context) (CustomClaims, error) {
	ginContext, ok := ctx.(*gin.Context)
	if !ok {
		return CustomClaims{}, failedToExtractTokenClaims
	}

	claims, exists := ginContext.Get("claims")
	if !exists {
		return CustomClaims{}, failedToExtractTokenClaims
	}

	tokenClaims, ok := claims.(*CustomClaims)
	if !ok {
		return CustomClaims{}, failedToExtractTokenClaims
	}

	return *tokenClaims, nil
}
