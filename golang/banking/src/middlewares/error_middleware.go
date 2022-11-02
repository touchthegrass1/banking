package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ErrorMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			logger.Error("error catched by middleware", zap.Error(err))
			switch err.Err {
			case gorm.ErrRecordNotFound:
				c.JSON(http.StatusNotFound, err.Error())
			case gorm.ErrInvalidTransaction:
				c.JSON(http.StatusBadRequest, err.Error())
			case gorm.ErrInvalidData:
				c.JSON(http.StatusBadRequest, err.Error())
			}
		}

		c.JSON(-1, "")
	}
}
