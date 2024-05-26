package initialize

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/requestid"
	"github.com/sevenzx/eztodo/config"
	"github.com/sevenzx/eztodo/middleware"
	"github.com/sevenzx/eztodo/router"
)

func HertzStart() {
	h := server.Default(server.WithHostPorts(config.Config.Server.HostPort))
	h.Use(requestid.New(), middleware.AccessLog())
	baseRouter := h.Group(config.Config.Server.BaseRouter)
	router.UserRouter(baseRouter)
	h.Spin()
}
