package api

import (
	"context"
	"github.com/sevenzx/eztodo/util/jwt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/sevenzx/eztodo/core/v1/service"
	"github.com/sevenzx/eztodo/model"
	"github.com/sevenzx/eztodo/model/response"
)

type userApi struct{}

// Register 注册
func (api *userApi) Register(c context.Context, ctx *app.RequestContext) {
	var user model.User
	_ = ctx.BindJSON(&user)
	err := service.User.Register(&user)
	if err != nil {
		hlog.Error(err)
		response.FailWithMsg(ctx, err.Error())
	} else {
		response.Ok(ctx)
	}
}

// Login 登录
func (api *userApi) Login(c context.Context, ctx *app.RequestContext) {
	var u model.User
	_ = ctx.BindJSON(&u)
	user, err := service.User.Login(u.Username, u.Password)
	if err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}
	// 登录成功 签发jwt
	j := jwt.NewHelper()
	claims := j.CreateClaims(model.CustomClaims{
		UUID:     user.UUID,
		Username: user.Username,
		Nickname: user.Nickname,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		hlog.Error(err)
		response.FailWithMsg(ctx, err.Error())
		return
	}
	// 向客户端设置token
	jwt.SetToken(ctx, token, int(claims.ExpiresAt.Unix()-time.Now().Unix()))
	response.OkWithData(ctx, map[string]interface{}{
		"user":       user,
		"token":      token,
		"expires_at": claims.ExpiresAt.Format(time.DateTime),
	})
}

// Information 通过ctx获取登录用户的信息
func (api *userApi) Information(c context.Context, ctx *app.RequestContext) {
	claims := jwt.GetClaims(ctx)
	u, err := service.User.GetUserByUuid(claims.UUID)
	if err != nil {
		hlog.Error(err)
		response.FailWithMsg(ctx, err.Error())
	} else {
		response.OkWithData(ctx, u)
	}
}
