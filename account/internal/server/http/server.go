package http

import (
	"net/http"

	"account/internal/service"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

var (
	u *service.Service
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	u = s
	engine = bm.DefaultServer(hc.Server)
	//pb.RegisterAccountBMServer(engine, u)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/account")
	{
		g.GET("/info", Info)
		g.GET("/profile", Profile)
		g.POST("/add", AddInfo)
		g.POST("/email/update", SetEmail)
		g.POST("/phone/update", SetPhone)
	}
}

func ping(ctx *bm.Context) {
	if err := u.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
