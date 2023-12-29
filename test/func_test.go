package test

import (
	"eee/dao/mysql"
	"eee/settings"
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestRecentPros(t *testing.T) {
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err: %v\n", err)
		return
	}
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	for i := 0; i < 100; i++ {
		_, err := mysql.RecentPros()
		if err != nil {
			t.Errorf("test error %s\n", err.Error())
		}
	}
}

func TestRecentOutput(t *testing.T) {
	for i := 0; i < 100; i++ {
		_, err := mysql.RecentOutput()
		if err != nil {
			log.Panicf("test error %s\n", err.Error())
		}
	}
}

func TestUnit(t *testing.T) {
	var c float64 = 0
	err := settings.Init()
	if err != nil {
		return
	}
	err = mysql.Init(settings.Conf.MySQLConfig)
	if err != nil {
		return
	}
	defer mysql.Close()
	for i := 0; i < 100; i++ {
		for j := 0; j < 10; j++ {
			_, err := mysql.GetProList(strconv.Itoa(j), strconv.Itoa(i))
			if err != nil {
				c++
			}
		}
	}
	t.Logf("percentage of the passed is: %f%%", ((1000-c)/1000)*100)
}
