package controller

import (
	"eee/dao/mysql"
	"eee/model"
	"eee/pkg/code"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type updatePro struct {
}

func UpdatePro(c *gin.Context) {
	pro := new(model.Project)
	c.ShouldBind(pro)
	log.Println(pro)
	pro.StartTime = strings.Split(pro.StartTime, "T")[0]
	pro.EndTime = strings.Split(pro.EndTime, "T")[0]
	err := mysql.UpdateP(pro)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}

func UpdateOPro(c *gin.Context) {
	id := c.Query("projectID")
	finish := c.Query("finish")
	log.Println(id, finish)
	if err := mysql.SetPFinish(id, finish); err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}
