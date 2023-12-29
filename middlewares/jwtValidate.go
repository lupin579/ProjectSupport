package middlewares

import (
	"eee/controller"
	"eee/pkg/code"
	"eee/pkg/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

func ValidateToken(c *gin.Context) {
	token := c.GetHeader("Authentic")
	if token == "" {
		controller.ResponseError(c, code.WithoutToken, errors.New("未携带token"))
		c.Abort()
		return
	}
	if _, err := utils.JWTValidator(token); err != nil {
		controller.ResponseError(c, code.InvalidToken, err)
		c.Abort()
		return
	}
	c.Next()
}
