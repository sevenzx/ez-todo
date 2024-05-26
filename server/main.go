package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gofrs/uuid/v5"
	"github.com/pkg/errors"
	"github.com/sevenzx/eztodo/conf"
	"github.com/sevenzx/eztodo/global"
	"github.com/sevenzx/eztodo/initialize"
	"github.com/sevenzx/eztodo/model"
	"github.com/sevenzx/eztodo/util"
	"gorm.io/gorm"
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

	h.POST("/user/register", func(c context.Context, ctx *app.RequestContext) {
		var user model.User
		_ = ctx.BindJSON(&user)
		if !errors.Is(global.DB.Where("username = ?", user.Username).First(&user).Error, gorm.ErrRecordNotFound) {
			ctx.JSON(consts.StatusOK, utils.H{"data": false})
		} else {
			user.UUID = uuid.Must(uuid.NewV4())
			user.Password = util.BcryptHash(user.Password)
			err := global.DB.Create(&user).Error
			if err != nil {
				ctx.JSON(consts.StatusOK, utils.H{"data": false})
				hlog.Error(err)
			} else {
				ctx.JSON(consts.StatusOK, utils.H{"data": user})
			}
		}
	})

	h.POST("/user/get", func(c context.Context, ctx *app.RequestContext) {
		var user model.User
		_ = ctx.BindJSON(&user)
		if errors.Is(global.DB.Where("id = ?", user.Id).First(&user).Error, gorm.ErrRecordNotFound) {
			ctx.JSON(consts.StatusOK, utils.H{"data": false})
		} else {
			ctx.JSON(consts.StatusOK, utils.H{"data": user})
		}
	})

	h.Spin()
}
