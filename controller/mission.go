package controller

import (
	"eee/dao/mysql"
	"eee/pkg/code"
	"eee/service"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type mfSerial struct {
	ID        int    `json:"id"`
	Uname     string `json:"uname"`
	Mid       int    `json:"m_id"`
	Mname     string `json:"m_name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Tag       string `json:"tag"`
	Finish    string `json:"finish"`
}

type FinalMission struct {
	ID        int        `json:"id"`
	Uname     string     `json:"uname"`
	Mid       int        `json:"m_id"`
	Mname     string     `json:"m_name"`
	StartTime string     `json:"startTime"`
	EndTime   string     `json:"endTime"`
	Tag       string     `json:"tag"`
	Finish    string     `json:"finish"`
	MfList    []mfSerial `json:"children"`
}

func Mission(c *gin.Context) {
	uname := c.Param("uname")
	pagenum := c.Query("pagenum")
	pagesize := c.Query("pagesize")
	log.Println(uname)
	var missionService service.MissionService
	list, err := missionService.MissionList(uname, pagenum, pagesize)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, list)
}

func GetMP(c *gin.Context) {
	pro := c.Param("pro")
	log.Println(pro)
	mfList, err := mysql.GetMP(pro)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	mfMap := make(map[string][]mfSerial)
	for _, mf := range mfList {
		mfs := mfSerial{
			ID:        mf.Mid,
			Uname:     mf.Uname,
			Mid:       mf.Mid,
			Mname:     mf.Mname,
			StartTime: strings.Split(mf.StartTime, "T")[0],
			EndTime:   strings.Split(mf.EndTime, "T")[0],
			Tag:       mf.Tag,
			Finish:    mf.Finish,
		}
		mfMap[mf.Uname] = append(mfMap[mf.Uname], mfs)
	}
	realCount := 1
	var finalMList []FinalMission
	for _, mList := range mfMap {
		var fm FinalMission
		for count, m := range mList {
			if count == 0 {
				fm.ID = realCount
				fm.Mname = m.Mname
				fm.EndTime = m.EndTime
				fm.StartTime = m.StartTime
				fm.Tag = m.Tag
				fm.Mid = m.Mid
				fm.Finish = m.Finish
				fm.Uname = m.Uname
			} else {
				mfs := mfSerial{
					ID:        realCount,
					Mname:     m.Mname,
					Uname:     m.Uname,
					Mid:       m.Mid,
					StartTime: m.StartTime,
					EndTime:   m.EndTime,
					Tag:       m.Tag,
					Finish:    m.Finish,
				}
				fm.MfList = append(fm.MfList, mfs)
			}
			realCount++
		}
		finalMList = append(finalMList, fm)
	}
	ResponseSuccess(c, code.Success, finalMList)
}
