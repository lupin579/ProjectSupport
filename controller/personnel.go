package controller

import (
	"eee/dao/mysql"
	"eee/model"
	"eee/pkg/code"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type p struct {
	UserList []*model.UserExtra
	Cur      string
	Total    string
}

func PeopleList(c *gin.Context) {
	pagenum := c.Query("pagenum")
	pagesize := c.Query("pagesize")
	pro := c.Query("pro")
	log.Println(pro)
	userList, err := mysql.UserList(pagenum, pagesize, pro)

	if err != nil {
		ResponseError(c, code.GetError, err)
		return
	}
	for _, user := range userList {
		onBoarding, _ := strconv.Atoi(strings.Split(user.Onboarding, "-")[0])
		fmt.Println(onBoarding)
		cur := time.Now().Year()
		user.WorkingTime = cur - onBoarding
		user.Onboarding = strings.Split(user.Onboarding, "T")[0]
	}
	total, err := mysql.CountH()
	if err != nil {
		ResponseError(c, code.ServerBusy, err)
		return
	}
	pl := p{
		UserList: userList,
		Cur:      pagenum,
		Total:    *total,
	}
	ResponseSuccess(c, code.Success, pl)
}

func FreePersonList(c *gin.Context) {
	log.Println("FreePersonList")
	pro := c.Query("pro")
	list, err := mysql.GetFreeWorker(pro)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	for _, fp := range list {
		onTime := strings.Split(fp.Onboarding, "-")[0]
		nowTime := time.Now().Year()
		oTime, _ := strconv.Atoi(onTime)
		fp.WorkingTime = nowTime - oTime
		fp.Onboarding = strings.Split(fp.Onboarding, "T")[0]
	}
	ResponseSuccess(c, code.Success, list)
}

func PersonList(c *gin.Context) {
	pagenum := c.Query("pagenum")
	pagesize := c.Query("pagesize")
	log.Println("PersonList")
	userList, err := mysql.GetAll(pagenum, pagesize)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
	}
	for _, user := range userList {
		onBoarding, _ := strconv.Atoi(strings.Split(user.Onboarding, "-")[0])
		fmt.Println(onBoarding)
		cur := time.Now().Year()
		user.WorkingTime = cur - onBoarding
		user.Onboarding = strings.Split(user.Onboarding, "T")[0]
	}
	total, err := mysql.CountH()
	if err != nil {
		ResponseError(c, code.ServerBusy, err)
		return
	}
	pl := p{
		UserList: userList,
		Cur:      pagenum,
		Total:    *total,
	}
	ResponseSuccess(c, code.Success, pl)
}
