package main

import (
	"eee/cache"
	"eee/dao/mysql"
	"eee/routers"
	"eee/settings"
	"fmt"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err: %v\n", err)
		return
	}
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := cache.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
	}
	defer cache.Close()
	r := routers.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf("0.0.0.0:%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
