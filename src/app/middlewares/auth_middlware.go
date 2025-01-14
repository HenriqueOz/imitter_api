package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	blackListService := services.NewBlackListService()
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

		token := utils.ParseToken(splitTokenString[1])
		if token == nil {
			c.JSON(http.StatusUnauthorized, utils.ResponseError(
				apperrors.ErrInvalidToken,
				apperrors.ErrInvalidClaims.Error(),
			))
			c.Abort()
			return
		}

		var uuid string
		claims := token.Claims.(jwt.MapClaims)
		if c.Request.URL.Path == refreshPath {
			uuid = claims["uuid"].(string)
			jti := claims["jti"].(string)

			err := blackListService.AddTokenToBlacklist(jti)
			if err != nil {
				c.JSON(http.StatusUnauthorized, utils.ResponseError(
					apperrors.ErrInvalidToken,
					err.Error(),
				))
				c.Abort()
				return
			}
		} else {
			uuid = claims["sub"].(string)
		}

		c.Request.Header.Add("uuid", uuid)

		c.Next()
	}
}
