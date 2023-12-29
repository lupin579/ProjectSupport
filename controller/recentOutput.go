package controller

import (
	"eee/dao/mysql"
	"eee/pkg/code"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Money struct {
	ProjectID      int     `json:"p_id"`
	Leader         string  `json:"leader"`
	StartTime      string  `json:"start_time"`
	EndTime        string  `json:"end_time"`
	TotalM         int     `json:"totalOutput"`    //总产值
	Amount         int     `json:"amount"`         //合同额
	Persons        int     `json:"personPercent"`  //人力预算比例
	PCost          int     `json:"personCost"`     //人力成本
	Balance        int     `json:"balance"`        //尾款
	BalancePercent float64 `json:"balancePercent"` //尾款占比/合同额
	OutRisk        int     `json:"outRisk"`        //外部风险
	InRisk         int     `json:"inRisk"`         //内部风险
	TotalCost      int     `json:"totalCost"`      //总成本
	ProfitPercent  float64 `json:"profitPercent"`  //利润率
}

type ROP struct {
	List  []int `json:"profit"`
	Total int   `json:"total"`
}

func RecentOutput(c *gin.Context) {
	log.Println("RecentOutput")
	list, err := mysql.RecentOutput()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	profit := make([]int, 7)
	for i, m := range list {
		fmt.Println(m)
		profit[i] = m.Amount - m.Input + m.Balance
	}
	ResponseSuccess(c, code.Success, profit)
}

func RecentPros(c *gin.Context) {
	log.Println("RecentPros")
	list, err := mysql.RecentPros()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	for _, p := range list {
		p.StartTime = strings.Split(p.StartTime, "T")[0]
		p.EndTime = strings.Split(p.EndTime, "T")[0]
	}
	ResponseSuccess(c, code.Success, list)
}

func ProMoney(c *gin.Context) {
	p_id := c.Param("pro")
	log.Println(p_id)
	proMoney, err := mysql.ProMoney(p_id)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := 20 + r.Intn(20)
	profitPercent, err := strconv.ParseFloat(fmt.Sprintf("%.2f", (float64(proMoney.Balance)+float64(proMoney.Amount)-float64(proMoney.Input)-20)/(float64(proMoney.Balance)+float64(proMoney.Amount))), 64)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	balancePercent, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(proMoney.Balance)/float64(proMoney.Amount)), 64)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	money := Money{
		ProjectID:      proMoney.ProjectID,
		Leader:         proMoney.Leader,
		StartTime:      strings.Split(proMoney.StartTime, "T")[0],
		EndTime:        strings.Split(proMoney.EndTime, "T")[0],
		TotalM:         proMoney.Amount + proMoney.Balance,
		Amount:         proMoney.Amount,
		Persons:        randNum,
		PCost:          proMoney.Input,
		Balance:        proMoney.Balance,
		BalancePercent: balancePercent,
		InRisk:         3 + randNum,
		OutRisk:        18 + randNum,
		TotalCost:      proMoney.Input + 20,
		ProfitPercent:  profitPercent,
	}
	ResponseSuccess(c, code.Success, money)
}
