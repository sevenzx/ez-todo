package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	cts "github.com/sevenzx/eztodo/common/consts"
)

type Response struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	RequestId string      `json:"request_id"`
}

func Result(ctx *app.RequestContext, code int, data interface{}, msg string) {
	requestId := ctx.Response.Header.Get(cts.RequestIdHeaderKey)
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
	Result(ctx, Success.Code, data, Success.Message)
}

func Fail(ctx *app.RequestContext) {
	FailWithError(ctx, ErrOperate)
}

func FailWithError(ctx *app.RequestContext, err Error) {
	Result(ctx, err.Code, nil, err.Message)
}

func FailWithMsg(ctx *app.RequestContext, msg string) {
	Result(ctx, ErrOperate.Code, nil, msg)
}

func UnAuth(ctx *app.RequestContext) {
	requestId := ctx.Response.Header.Get(cts.RequestIdHeaderKey)
	ctx.JSON(consts.StatusUnauthorized, Response{
		Code:      ErrUnAuth.Code,
		Data:      nil,
		Msg:       ErrUnAuth.Message,
		RequestId: requestId,
	})
}
