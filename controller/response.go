package controller

import (
	"eee/pkg/code"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(ctx *gin.Context, errCode int, data interface{}) {
	ctx.JSON(200, Response{
		Code:    errCode,
		Message: code.Msg(errCode),
		Data:    data,
	})
}

func ResponseError(ctx *gin.Context, errCode int, err error) {
	ctx.JSON(500, Response{
		Code:    errCode,
		Message: code.Msg(errCode),
		Data:    err.Error(),
	})
}
