package service

import (
	"eee/dao/mysql"
	"eee/model"
	"strings"
	"time"
)

type UpdateMissionService struct {
	MissionId   uint16 `json:"m_id"`
	MissionName string `json:"m_name"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Tag         string `json:"tag"`
	Leader      string `json:"leader"`
	ProjectId   uint16 `json:"project"`
	Description string `json:"description"`
}

func (updateMissionService *UpdateMissionService) UpdateMission() (err error) {
	mission := new(model.Mission)
	mission.MissionId = updateMissionService.MissionId
	mission.MissionName = updateMissionService.MissionName
	mission.StartTime = strings.Split(updateMissionService.StartTime, "T")[0]
	mission.EndTime = strings.Split(updateMissionService.EndTime, "T")[0]
	s, err := time.Parse("2006-01-02", mission.StartTime)
	if err != nil {
		return
	}
	st := s.AddDate(0, 0, 1)
	mission.StartTime = st.Format("2006-01-02")
	e, err := time.Parse("2006-01-02", mission.EndTime)
	if err != nil {
		return
	}
	en := e.AddDate(0, 0, 1)
	mission.EndTime = en.Format("2006-01-02")
	mission.Tag = updateMissionService.Tag
	mission.Leader = updateMissionService.Leader
	mission.ProjectId = updateMissionService.ProjectId
	mission.Description = updateMissionService.Description
	err = mysql.UpdateMission(mission)
	return
}
