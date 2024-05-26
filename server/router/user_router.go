package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	v1 "github.com/sevenzx/eztodo/api/v1"
)

func UserRouter(r *route.RouterGroup) {
	group := r.Group("/user")

	api := v1.UserApi{}
	group.POST("/register", api.Register)
	group.POST("/get", api.GetById)
	group.GET("/get/:id", api.GetById)
}
