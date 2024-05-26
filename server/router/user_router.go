package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/sevenzx/eztodo/api/v1"
)

func UserRouter(r *route.RouterGroup) {
	group := r.Group("/user")

	api := v1.API.User
	group.POST("/register", api.Register)
	group.POST("/get", api.GetById)
}
