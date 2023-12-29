package controller

import (
	"eee/pkg/code"
	"eee/service"
	"log"

	"github.com/gin-gonic/gin"
)

func UpdateMission(c *gin.Context) {
	updateMissionService := new(service.UpdateMissionService)
	c.ShouldBind(updateMissionService)
	log.Println(updateMissionService)
	err := updateMissionService.UpdateMission()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}
