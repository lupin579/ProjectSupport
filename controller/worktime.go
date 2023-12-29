package controller

import (
	"eee/dao/mysql"
	"eee/pkg/code"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type wt struct {
	FirstS  map[int]string `json:"firsts"`
	FirstE  map[int]string `json:"firste"`
	SecondS map[int]string `json:"seconds"`
	SecondE map[int]string `json:"seconde"`
	ThirdS  map[int]string `json:"thirds"`
	ThirdE  map[int]string `json:"thirde"`
}

type res struct {
	UserName []string `json:"user_name"`
	FirstS   []string `json:"firsts"`
	FirstE   []string `json:"firste"`
	SecondS  []string `json:"seconds"`
	SecondE  []string `json:"seconde"`
	ThirdS   []string `json:"thirds"`
	ThirdE   []string `json:"thirde"`
}

func PostWorktime(c *gin.Context) {
	worktime := c.Query("worktime")
	m_id := c.Query("m_id")
	extraTime, err := mysql.GetWorkTime(m_id)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	n_worktime, _ := strconv.Atoi(worktime)
	fmt.Println("workTime:", n_worktime)
	n_extraTime, _ := strconv.Atoi(extraTime)
	fmt.Println("extraTime:", n_extraTime)
	realTime := n_extraTime + n_worktime
	fmt.Println("realTime:", realTime)
	err = mysql.PostWorkTime(realTime, m_id)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}

func MemberPercent(c *gin.Context) {
	pro := c.Query("pro")
	log.Println(pro)
	timeList, err := mysql.TimePercent(pro)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, timeList)
}

func TimeChart(c *gin.Context) {
	log.Println("TimeChart")
	cList, err := mysql.TimeChart()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	for _, c := range cList {
		fmt.Println(c.Uid)
	}
	num, err := mysql.CountT()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	nameList, err := mysql.GetUserName()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	w := wt{
		FirstS:  make(map[int]string, num),
		FirstE:  make(map[int]string, num),
		SecondS: make(map[int]string, num),
		SecondE: make(map[int]string, num),
		ThirdS:  make(map[int]string, num),
		ThirdE:  make(map[int]string, num),
	}
	for _, c := range cList {
		if c.ProID != nil {
			if _, notnull := w.FirstS[c.Uid]; !notnull {
				w.FirstS[c.Uid] = strings.Split((*c.StartTime), "T")[0]
				w.FirstE[c.Uid] = strings.Split((*c.EndTime), "T")[0]
			} else if _, notnull := w.SecondS[c.Uid]; !notnull {
				w.SecondS[c.Uid] = strings.Split((*c.StartTime), "T")[0]
				w.SecondE[c.Uid] = strings.Split((*c.EndTime), "T")[0]
			} else {
				w.ThirdS[c.Uid] = strings.Split((*c.StartTime), "T")[0]
				w.ThirdE[c.Uid] = strings.Split((*c.EndTime), "T")[0]
			}
		} else {
			w.FirstS[c.Uid] = "0"
			w.FirstE[c.Uid] = "0"
			w.SecondS[c.Uid] = "0"
			w.SecondE[c.Uid] = "0"
			w.ThirdS[c.Uid] = "0"
			w.ThirdE[c.Uid] = "0"
		}
	}
	for _, c := range cList {
		if _, notnull := w.SecondS[c.Uid]; !notnull {
			w.SecondS[c.Uid] = "0"
			w.SecondE[c.Uid] = "0"
			w.ThirdS[c.Uid] = "0"
			w.ThirdE[c.Uid] = "0"
		} else if _, notnull := w.ThirdS[c.Uid]; !notnull {
			w.ThirdS[c.Uid] = "0"
			w.ThirdE[c.Uid] = "0"
		}
	}
	var ids []int
	for id, _ := range w.FirstE {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	var r res
	for _, i := range ids {
		r.FirstS = append(r.FirstS, w.FirstS[i])
		r.FirstE = append(r.FirstE, w.FirstE[i])
		r.SecondS = append(r.SecondS, w.SecondS[i])
		r.SecondE = append(r.SecondE, w.SecondE[i])
		r.ThirdS = append(r.ThirdS, w.ThirdS[i])
		r.ThirdE = append(r.ThirdE, w.ThirdE[i])
	}
	for _, name := range nameList {
		r.UserName = append(r.UserName, name)
	}
	ResponseSuccess(c, code.Success, r)
}
