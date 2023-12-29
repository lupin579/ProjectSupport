package controller

import (
	"eee/dao/mysql"
	"eee/pkg/code"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	Uid        int     `json:"uid"`
	Uname      string  `json:"uname" `
	WorkType   string  `json:"work_type"`
	Onboarding string  `json:"noboarding"`
	Mobile     string  `json:"mobile" `
	Email      string  `json:"email" `
	Password   string  `json:"password" `
	Level      float64 `json:"level"`
}

func AddUser(c *gin.Context) {
	userService := new(UserService)
	c.ShouldBindJSON(userService)
	log.Println(userService)
	sid := strconv.Itoa(userService.Uid)
	err := mysql.UserAdd(sid, userService.Uname, userService.WorkType, userService.Onboarding, userService.Mobile, userService.Email, userService.Password, userService.Level)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}
