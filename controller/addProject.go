package controller

import (
	"eee/pkg/code"
	"eee/service"
	"log"

	"github.com/gin-gonic/gin"
)

func AddProject(c *gin.Context) {
	uname := c.Param("uname")
	addService := new(service.AddProService)
	err := c.ShouldBind(addService)
	log.Println(addService)
	if err != nil {
		ResponseError(c, code.BindError, err)
		return
	}
	addService.Uname = uname
	err = addService.AddPro()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}
