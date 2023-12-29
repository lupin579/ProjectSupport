package service

import (
	"eee/dao/mysql"
	"eee/model"
	"strings"
)

type ProListService struct {
	Uname    string
	Pagenum  string
	Pagesize string
}

type ProListSerial struct {
	List  []*model.Project
	Total string
	Cur   string
}

func (proListService *ProListService) ProListService() (resList []*model.Project, s *ProListSerial, err error) {
	right, err := mysql.GetRightByName(proListService.Uname)
	if err != nil {
		return nil, nil, err
	}
	proList, err := mysql.GetProList(proListService.Pagenum, proListService.Pagesize)
	if err != nil {
		return proList, nil, err
	}
	for _, pro := range proList {
		pro.StartTime = strings.Split(pro.StartTime, "T")[0]
		pro.EndTime = strings.Split(pro.EndTime, "T")[0]
	}
	num, err := mysql.CountP()
	if *right == "开发者" {
		for _, pro := range proList {
			pro.Balance = 0
			pro.Input = 0
		}
	}
	proListSerial := ProListSerial{
		List:  proList,
		Cur:   proListService.Pagenum,
		Total: *num,
	}
	return nil, &proListSerial, err
}
