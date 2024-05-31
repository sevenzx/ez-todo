package middleware

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/sevenzx/eztodo/model/response"
)

func AccessLog() app.HandlerFunc {
	// 自定义一个requestId的标签
	accesslog.Tags["requestId"] = func(buf accesslog.Buffer, ctx *app.RequestContext, data *accesslog.Data, extraParam string) (int, error) {
		return buf.WriteString(ctx.Response.Header.Get(response.RequestIdHeaderKey))
	}
	return accesslog.New(
		// 自定义日志格式
		accesslog.WithFormat("[${requestId}] ${status} ${method} ${path} ${latency}"),
	)
}
