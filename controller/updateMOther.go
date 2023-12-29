package controller

import (
	"eee/dao/mysql"
	"eee/pkg/code"
	"log"

	"github.com/gin-gonic/gin"
)

func UpdateMFinish(c *gin.Context) {
	id := c.Query("missionID")
	finish := c.Query("finish")
	log.Println(id, finish)
	if err := mysql.SetFinish(id, finish); err != nil {
		ResponseError(c, code.OperateFail, err)
	}
	ResponseSuccess(c, code.Success, nil)
}
