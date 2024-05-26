package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
