package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/utils"
)

// TODO revisar o método de pegar o uuid de dentro do refresh token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.ResponseError(
				apperrors.ErrMissingHeaders,
				apperrors.ErrMissingAuthorization.Error(),
			))
			c.Abort()
			return
		}

		splitTokenString := strings.Split(authHeader, " ")
		if len(splitTokenString) < 2 || strings.Compare(splitTokenString[0], "Bearer") != 0 {
			c.JSON(http.StatusUnauthorized, utils.ResponseError(
				apperrors.ErrInvalidToken,
				apperrors.ErrTokenFormat.Error(),
			))
			c.Abort()
			return
		}

		token := parseToken(splitTokenString[1])
		if token != nil {
			claims := token.Claims.(jwt.MapClaims)
			if c.Request.URL.Opaque == "/refresh" {
				claims = getTokenClaimsUnverified(claims["sub"].(string))
			}
			c.Header("uuid", claims["sub"].(string))
		}

		c.Next()
	}
}

func getTokenClaimsUnverified(expiredToken string) jwt.MapClaims {
	token, _, err := new(jwt.Parser).ParseUnverified(expiredToken, jwt.MapClaims{})
	if err != nil {
		log.Printf("Failed parsing token: %v\n", err)
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		if _, exists := claims["sub"].(string); !exists {
			return nil
		}
		return claims
	}

	return nil
}

func parseToken(tokenString string) *jwt.Token {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTSECRET")), nil
	})

	if err != nil {
		log.Printf("Failed parsing token: %v\n", err)
		return nil
	}
	return token
}
