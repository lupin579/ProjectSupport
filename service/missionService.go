package service

import (
	"eee/dao/mysql"
	"eee/model"
	"strconv"
	"strings"
)

type MissionService struct {
	Missions []*model.Mission
	Total    int
	Cur      int
}

func (missionService *MissionService) MissionList(uname, pagenum, pagesize string) (missionRes *MissionService, err error) {
	num, _ := strconv.Atoi(pagenum)
	pagenum = strconv.Itoa(num - 1)
	missionList, err := mysql.MissionList(uname, pagenum, pagesize)
	if err != nil {
		return
	}
	for _, mission := range missionList {
		mission.StartTime = strings.Split(mission.StartTime, "T")[0]
		mission.EndTime = strings.Split(mission.EndTime, "T")[0]
	}
	missionRes = new(MissionService)
	missionRes.Missions = missionList
	missionRes.Cur = num
	total, err := mysql.Count(uname)
	if err != nil {
		return
	}
	missionRes.Total, _ = strconv.Atoi(*total)
	return
}
