package controller

import (
	"eee/pkg/code"
	"eee/service"
	"log"

	"github.com/gin-gonic/gin"
)

func ChangePassword(c *gin.Context) {
	changePassword := new(service.ChangePassword)
	changePassword.Uname = c.Param("uname")
	if err := changePassword.ChangePasswordService(); err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	log.Println(changePassword)
	ResponseSuccess(c, code.Success, nil)
}
