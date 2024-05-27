package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/sevenzx/eztodo/core/v1/api"
)

func registerUserRouter(r *route.RouterGroup) {
	group := r.Group("/user")

	group.POST("/register", api.User.Register)
	group.POST("/get", api.User.GetById)
}
