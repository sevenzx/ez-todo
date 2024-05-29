package router

import "github.com/cloudwego/hertz/pkg/route"

// RegisterV1Router 注册V1接口的路由
func RegisterV1Router(r *route.RouterGroup) {
	v1Router := r.Group("/v1")

	registerUserRouter(v1Router)
}
