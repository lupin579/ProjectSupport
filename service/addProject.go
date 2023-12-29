package service

import (
	"eee/dao/mysql"
	"eee/model"
	"errors"
)

type AddProService struct {
	Uname string
	model.Project
}

func (addProService *AddProService) AddPro() (err error) {
	right, err := mysql.GetRightByName(addProService.Uname)
	if err != nil {
		return
	}
	if *right != "项目经理" {
		return errors.New("权限错误，只有项目经理才能创建项目")
	}
	var pro model.Project = addProService.Project
	pro.Finish = "否"
	err = mysql.AddProject(pro)
	return
}
