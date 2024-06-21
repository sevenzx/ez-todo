package middleware

import (
	"context"
	"github.com/sevenzx/eztodo/common/consts"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/sevenzx/eztodo/config"
	"github.com/sevenzx/eztodo/model/response"
	"github.com/sevenzx/eztodo/util"
	jwtutil "github.com/sevenzx/eztodo/util/jwt"
	"time"
)

func JWTAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时会返回token信息
		// 这里前端需要把token存储到cookie或者本地localStorage中
		token := jwtutil.GetToken(ctx)
		// 1. 判断是否有token
		if token == "" {
			response.UnAuth(ctx)
			ctx.Abort()
			return
		}
		// 2. 验证token
		j := jwtutil.NewHelper()
		claims, err := j.ParseToken(token)
		if err != nil {
			// 如果token过期就清除客户端的token
			if errors.Is(err, jwt.ErrTokenExpired) {
				jwtutil.ClearToken(ctx)
			}
			response.FailWithMsg(ctx, err.Error())
			ctx.Abort()
			return
		}
		// 3. 在上下文中设置claims供后续使用
		ctx.Set(consts.JwtClaimsKey, claims)
		// 4. 判断是否需要刷新
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			duration, _ := util.ParseDuration(config.Config.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))
			newToken, _ := j.RefreshToken(token, *claims)
			ctx.Header("new-token", newToken)
			jwtutil.SetToken(ctx, newToken, int(duration.Seconds()))
		}
		ctx.Next(c)
	}
}
