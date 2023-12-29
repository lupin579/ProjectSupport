package service

import (
	"eee/dao/mysql"
	"errors"
	"fmt"
)

type RegisterService struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Right    string `json:"right"`
}

func (registerService *RegisterService) Register() (err error) {
	if exist, _ := mysql.IsExist(registerService.UserName); !exist {
		fmt.Println(exist)
		return errors.New("用户不存在")
	}

	if realPassword, err := mysql.GetPassword(registerService.UserName); err != nil {
		return err
	} else if *realPassword != registerService.Password {
		return errors.New("密码错误")
	}
	var realRight *string
	if realRight, err = mysql.GetRightByUserName(registerService.UserName); err != nil {
		return err
	} else if *realRight != registerService.Right {
		return errors.New("身份错误")
	}
	return nil
}
