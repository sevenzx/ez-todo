package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/sevenzx/eztodo/core/v1/api"
	"github.com/sevenzx/eztodo/middleware"
)

func registerUserRouter(r *route.RouterGroup) {
	basicRouter := r.Group("/user")
	userRouter := r.Group("/user")
	userRouter.Use(middleware.JWTAuth())
	{
		// 注册登录不需要JWT验证
		basicRouter.POST("/register", api.User.Register)
		basicRouter.POST("/login", api.User.Login)
	}
	{
		userRouter.GET("/info", api.User.Information)
		userRouter.GET("/logout", api.User.Logout)
	}
}
