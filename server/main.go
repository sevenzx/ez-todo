package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sevenzx/eztodo/conf"
	"github.com/sevenzx/eztodo/global"
	"github.com/sevenzx/eztodo/initialize"
)

func main() {
	// 加载配置文件
	initialize.LoadConfigFile()
	// 初始化数据库
	global.DB = initialize.GormMysql()
	if global.DB != nil {
		// 初始化表
		initialize.RegisterTables()
		db, _ := global.DB.DB()
		defer db.Close()
	}

	h := server.Default(server.WithHostPorts(conf.Config.Server.HostPort))
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.Spin()
}
