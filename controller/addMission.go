package controller

import (
	"eee/pkg/code"
	"eee/service"
	"log"

	"github.com/gin-gonic/gin"
)

func AddMission(c *gin.Context) {
	addMissionService := new(service.AddMissionService)
	c.ShouldBind(addMissionService)
	log.Println(addMissionService)
	err := addMissionService.AddMission()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}
