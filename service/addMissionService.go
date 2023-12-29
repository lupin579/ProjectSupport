package service

import (
	"eee/dao/mysql"
	"eee/model"
	"fmt"
	"strings"
	"time"
)

type AddMissionService struct {
	MissionId   uint16 `json:"m_id"`
	MissionName string `json:"m_name"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Tag         string `json:"tag"`
	Leader      string `json:"leader"`
	ProjectId   uint16 `json:"project"`
	Description string `json:"description"`
	Worker      string `json:"worker"`
}

func (addMissionService *AddMissionService) AddMission() (err error) {
	fmt.Println(addMissionService)
	mission := new(model.Mission)
	mission.MissionId = addMissionService.MissionId
	mission.MissionName = addMissionService.MissionName
	StartTime := strings.Split(addMissionService.StartTime, "T")[0]
	realS, err := time.Parse("2006-01-02", StartTime)
	realS.AddDate(0, 0, 1)
	EndTime := strings.Split(addMissionService.EndTime, "T")[0]
	realE, _ := time.Parse("2006-01-02", EndTime)
	realE.AddDate(0, 0, 1)
	mission.StartTime = realS.Format("2006-01-02")
	mission.EndTime = realE.Format("2006-01-02")
	mission.Tag = addMissionService.Tag
	mission.Leader = addMissionService.Leader
	mission.ProjectId = addMissionService.ProjectId
	mission.Description = addMissionService.Description
	mission.Worker = addMissionService.Worker
	fmt.Println(mission)
	err = mysql.AddMission(mission)
	return
}
