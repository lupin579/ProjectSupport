package controller

import (
	"eee/dao/mysql"
	"eee/model"
	"eee/pkg/code"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ResPro struct {
	ProDetails *model.Project
	WorkerList []*model.WS
}

func ProDetails(c *gin.Context) {
	pro := c.Param("pro")
	log.Println(pro)
	proDet, err := mysql.ProDetails(pro)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	proDet.StartTime = strings.Split(proDet.StartTime, "T")[0]
	proDet.EndTime = strings.Split(proDet.EndTime, "T")[0]
	wList, err := mysql.WorkerList(pro)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	for _, ws := range wList {
		ws.StartTime = strings.Split(ws.StartTime, "T")[0]
		ws.EndTime = strings.Split(ws.EndTime, "T")[0]
		start, err := time.Parse("2006-01-02", ws.StartTime)
		if err != nil {
			ResponseError(c, code.OperateFail, err)
			return
		}
		end, err := time.Parse("2006-01-02", ws.EndTime)
		if err != nil {
			ResponseError(c, code.OperateFail, err)
			return
		}
		if start.Unix() <= time.Now().Unix() && time.Now().Unix() <= end.Unix() {
			ws.IsSupport = 1
		} else {
			ws.IsSupport = 0
		}
	}
	res := ResPro{
		ProDetails: &proDet,
		WorkerList: wList,
	}
	ResponseSuccess(c, code.Success, res)
}
