package auth

import (
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/metadata"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	authrpc "github.com/itswcg/micro/auth/api"
)

type Config struct {
	AuthClient *warden.ClientConfig
}

type Auth struct {
	config *Config
	ac     authrpc.AuthClient
}

func New(conf *Config) *Auth {
	auth := &Auth{config: conf}

	ac, err := authrpc.NewClient(auth.config.AuthClient)
	if err != nil {
		panic(err)
	}
	auth.ac = ac
	return auth
}

func (a *Auth) authToken(ctx *bm.Context, token string) (mid int64, err error) {
	tokenReq := &authrpc.TokenReq{Token: token}
	tokenReply, err := a.ac.Token(ctx, tokenReq)
	if err != nil {
		return 0, err
	}
	return tokenReply.Mid, nil
}

func (a *Auth) IsAdmin(ctx *bm.Context) {

}

func (a *Auth) IsAuthenticated(ctx *bm.Context) {
	req := ctx.Request
	token := req.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(nil, ecode.Unauthorized)
		ctx.Abort()
		return
	}
	mid, err := a.authToken(ctx, token)
	if err != nil {
		ctx.JSON(nil, ecode.Unauthorized)
		ctx.Abort()
		return
	}
	a.setMid(ctx, mid)
}

func (a *Auth) AllowAny(ctx *bm.Context) {

}

func (a *Auth) setMid(ctx *bm.Context, mid int64) {
	ctx.Set(metadata.Mid, mid)
	if md, ok := metadata.FromContext(ctx); ok {
		md[metadata.Mid] = mid
		return
	}
}
