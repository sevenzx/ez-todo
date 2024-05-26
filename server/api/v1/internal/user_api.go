package internal

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/sevenzx/eztodo/model"
	"github.com/sevenzx/eztodo/model/response"
	"github.com/sevenzx/eztodo/service"
)

type UserApi struct{}

func (api UserApi) Register(c context.Context, ctx *app.RequestContext) {
	var user model.User
	_ = ctx.BindJSON(&user)
	err := service.User.Register(&user)
	hlog.Info(err)
	if err != nil {
		hlog.Error(err)
		response.Result(ctx, -1, nil, err.Error())
	} else {
		response.Result(ctx, 0, user, "success")
	}
}

func (api UserApi) GetById(c context.Context, ctx *app.RequestContext) {
	var user model.User
	_ = ctx.BindJSON(&user)
	u, err := service.User.GetById(user.Id)
	if err != nil {
		hlog.Error(err)
		response.Result(ctx, -1, nil, err.Error())
	} else {
		response.Result(ctx, 0, u, "success")
	}
}
