package controller

import (
	"eee/dao/mysql"
	"eee/model"
	"eee/pkg/code"
	"log"

	"github.com/gin-gonic/gin"
)

func Distribute(c *gin.Context) {
	su := new(model.Support)
	c.ShouldBind(su)
	log.Println(su)
	err := mysql.Distribute(su)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}

func Cancel(c *gin.Context) {
	userName := c.Query("userName")
	proID := c.Query("proID")
	log.Println(userName, " ", proID)
	err := mysql.Cancel(userName, proID)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}
