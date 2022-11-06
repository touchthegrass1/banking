package middlewares

import (
	"net/http"

	"github.com/dopefresh/banking/golang/banking/src/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthMiddleware(logger *zap.Logger, service services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := service.VerifyToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Bad jwt")
			return
		}
		userId, exists := token.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, "User id somehow didn't appear in jwt")
			return
		}
		userIdFloat, ok := userId.(float64)
		if !ok {
			c.JSON(http.StatusInternalServerError, "User id can't be converted to int. Go to authorization service")
			return
		}
		c.Set("userId", int64(userIdFloat))

		c.Next()
	}
}
