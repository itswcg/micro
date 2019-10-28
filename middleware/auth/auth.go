package auth

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	accrpc "github.com/itswcg/micro/account/api"
)

type Config struct {
	IsAccount     bool
	AccountClient *warden.ClientConfig
}

type Auth struct {
	config *Config
	ac     accrpc.AccountClient
}

func New(conf *Config) *Auth {
	if conf == nil {
		conf = &Config{IsAccount: false, AccountClient: nil}
	}
	auth := &Auth{config: conf}
	if auth.config.IsAccount == false {
		ac, err := accrpc.NewClient(auth.config.AccountClient)
		if err != nil {
			panic(err)
		}
		auth.ac = ac
	}

	return auth
}

func (a *Auth) authToken(ctx *bm.Context) (mid int64, err error) {
	return
}

func (a *Auth) IsAdmin(ctx *bm.Context) (mid int64, err error) {
	return
}

func (a *Auth) IsAuthenticated(ctx *bm.Context) (mid int64, err error) {
	return
}

func (a *Auth) AllowAny(ctx *bm.Context) (mid int64, err error) {
	return
}
