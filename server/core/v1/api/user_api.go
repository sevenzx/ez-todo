package api

import (
	"context"

	"github.com/sevenzx/eztodo/core/v1/service"
	"github.com/sevenzx/eztodo/model"
	"github.com/sevenzx/eztodo/model/response"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type userApi struct{}

func (api *userApi) Register(c context.Context, ctx *app.RequestContext) {
	var user model.User
	_ = ctx.BindJSON(&user)
	err := service.User.Register(&user)
	hlog.Info(err)
	if err != nil {
		hlog.Error(err)
		response.FailWithMsg(ctx, err.Error())
	} else {
		response.Ok(ctx)
	}
}

func (api *userApi) GetById(c context.Context, ctx *app.RequestContext) {
	var user model.User
	_ = ctx.BindJSON(&user)
	u, err := service.User.GetById(user.Id)
	if err != nil {
		hlog.Error(err)
		response.FailWithMsg(ctx, err.Error())
	} else {
		response.OkWithData(ctx, u)
	}
}