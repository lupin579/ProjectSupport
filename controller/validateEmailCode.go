package controller

import (
	"eee/pkg/code"
	"eee/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func ValidateEmailCode(c *gin.Context) {
	changePassword := new(service.ChangePassword)
	err := c.ShouldBindJSON(changePassword)
	log.Println(changePassword)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	fmt.Println(changePassword)
	err = changePassword.EmailCodeValidator()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}
