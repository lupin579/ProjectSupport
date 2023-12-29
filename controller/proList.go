package controller

import (
	"eee/dao/mysql"
	"eee/pkg/code"
	"eee/service"
	"log"

	"github.com/gin-gonic/gin"
)

func ProList(c *gin.Context) {
	uname := c.Param("uname")
	pagenum := c.Query("pagenum")
	pagesize := c.Query("pagesize")
	proList := new(service.ProListService)
	proList.Uname = uname
	proList.Pagenum = pagenum
	proList.Pagesize = pagesize
	log.Println(proList)
	_, list, err := proList.ProListService()
	if err != nil {
		ResponseError(c, code.GetError, err)
		return
	}
	ResponseSuccess(c, code.Success, list)
}

func MyPro(c *gin.Context) {
	uid := c.Param("uid")
	list, err := mysql.GetProByUid(uid)
	if err != nil {
		ResponseError(c, code.GetError, err)
		return
	}
	ResponseSuccess(c, code.Success, list)
}

func MyProOne(c *gin.Context) {
	uid := c.Param("uname")
	list, err := mysql.GetProByLeader(uid)
	if err != nil {
		ResponseError(c, code.GetError, err)
		return
	}
	ResponseSuccess(c, code.Success, list)
}
