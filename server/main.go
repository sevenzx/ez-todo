package main

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/sevenzx/eztodo/core/global"
	"github.com/sevenzx/eztodo/initialize"
)

func main() {
	// 加载配置文件
	initialize.LoadConfigFile()
	lumberjack := initialize.Lumberjack()
	hlog.SetOutput(lumberjack)
	// 初始化数据库
	global.DB = initialize.GormMysql(lumberjack)
	if global.DB != nil {
		// 初始化表
		initialize.RegisterTables()
		db, _ := global.DB.DB()
		defer db.Close()
	}
	// 开启Hertz
	initialize.HertzStart()
}
