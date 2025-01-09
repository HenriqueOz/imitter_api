package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	const refreshPath string = "/v1/auth/refresh"

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
		if token == nil {
			c.JSON(http.StatusUnauthorized, utils.ResponseError(
				apperrors.ErrInvalidToken,
				apperrors.ErrInvalidClaims.Error(),
			))
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if c.Request.URL.Path == refreshPath {
			jti := claims["jti"].(string)

			err := services.StoreClaimUuid(jti)
			if err != nil {
				c.JSON(http.StatusUnauthorized, utils.ResponseError(
					apperrors.ErrInvalidToken,
					err.Error(),
				))
				c.Abort()
				return
			}
		}

		uuid := claims["uuid"].(string)

		c.Request.Header.Add("uuid", uuid)

		c.Next()
	}
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
