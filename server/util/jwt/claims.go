package jwt

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/sevenzx/eztodo/model"
	"net"
)

const (
	TokenKey  = "x-token"
	ClaimsKey = "claims"
)

// SetToken 设置Token
func SetToken(ctx *app.RequestContext, token string, maxAge int) {
	// 增加cookie x-token
	host, _, err := net.SplitHostPort(string(ctx.Host()))
	if err != nil {
		host = string(ctx.Host())
	}

	if net.ParseIP(host) != nil {
		// "/"：cookie的路径，表示cookie在整个域中都有效
		// ""：cookie的域名，留空表示只在当前域有效
		ctx.SetCookie(TokenKey, token, maxAge, "/", "", protocol.CookieSameSiteDisabled, false, false)
	} else {
		ctx.SetCookie(TokenKey, token, maxAge, "/", host, protocol.CookieSameSiteDisabled, false, false)
	}
}

// GetToken 获取Token
func GetToken(ctx *app.RequestContext) string {
	bs := ctx.Cookie(TokenKey)
	token := string(bs)
	if token == "" {
		hbs := ctx.GetHeader(TokenKey)
		token = string(hbs)
	}
	return token
}

// ClearToken 清除Token
func ClearToken(ctx *app.RequestContext) {
	host, _, err := net.SplitHostPort(string(ctx.Host()))
	if err != nil {
		host = string(ctx.Host())
	}
	if net.ParseIP(host) != nil {
		ctx.SetCookie(TokenKey, "", -1, "/", "", protocol.CookieSameSiteDisabled, false, false)
	} else {
		ctx.SetCookie(TokenKey, "", -1, "/", host, protocol.CookieSameSiteDisabled, false, false)
	}
}

// GetClaims 获取Claims
func GetClaims(ctx *app.RequestContext) *model.Claims {
	value, exists := ctx.Get(ClaimsKey)
	if !exists {
		claims, err := GetClaimsFormToken(ctx)
		if err != nil {
			return nil
		} else {
			return claims
		}
	}
	claims, ok := value.(*model.Claims)
	if !ok {
		return nil
	} else {
		return claims
	}
}

// GetClaimsFormToken 从token中获取claims
func GetClaimsFormToken(ctx *app.RequestContext) (*model.Claims, error) {
	token := GetToken(ctx)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		return nil, err
	} else {
		return claims, nil
	}
}
