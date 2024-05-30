package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sevenzx/eztodo/model/response/exception"
)

const RequestIdHeaderKey = "X-Request-ID"

type Response struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	RequestId string      `json:"request_id"`
}

func Result(ctx *app.RequestContext, code int, data interface{}, msg string) {
	requestId := ctx.Response.Header.Get(RequestIdHeaderKey)
	ctx.JSON(consts.StatusOK, Response{
		Code:      code,
		Data:      data,
		Msg:       msg,
		RequestId: requestId,
	})
}

func Ok(ctx *app.RequestContext) {
	OkWithData(ctx, true)
}

func OkWithData(ctx *app.RequestContext, data interface{}) {
	Result(ctx, exception.Success.Code, data, exception.Success.Message)
}

func Fail(ctx *app.RequestContext) {
	FailWithException(ctx, exception.Operate)
}

func FailWithException(ctx *app.RequestContext, ex exception.Exception) {
	Result(ctx, ex.Code, nil, ex.Message)
}

func FailWithMsg(ctx *app.RequestContext, msg string) {
	Result(ctx, exception.Operate.Code, nil, msg)
}

func UnAuth(ctx *app.RequestContext) {
	requestId := ctx.Response.Header.Get(RequestIdHeaderKey)
	ctx.JSON(consts.StatusUnauthorized, Response{
		Code:      exception.UnAuth.Code,
		Data:      nil,
		Msg:       exception.UnAuth.Message,
		RequestId: requestId,
	})
}
