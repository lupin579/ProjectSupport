package controller

import (
	"eee/pkg/code"
	"eee/pkg/utils"
	"eee/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	registerService := new(service.RegisterService)
	c.ShouldBindJSON(&registerService)
	err := registerService.Register()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	token, err := utils.JWTIssue(registerService.UserName)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	c.Header("Authorization", token)
	ResponseSuccess(c, code.Success, nil)
}
