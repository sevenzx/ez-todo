package initialize

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/sevenzx/eztodo/config"
	"github.com/sevenzx/eztodo/core/v1/router"
	"github.com/sevenzx/eztodo/middleware"
)

func HertzStart() {
	h := server.Default(server.WithHostPorts(config.Config.Server.HostPort))
	h.Use(middleware.RequestId(), middleware.AccessLog())
	baseRouter := h.Group(config.Config.Server.BaseRouter)
	router.RegisterV1Router(baseRouter)
	h.Spin()
}
