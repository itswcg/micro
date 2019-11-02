package http

import (
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/itswcg/micro/middleware/auth"
	"net/http"

	"github.com/itswcg/micro/account/internal/service"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

var (
	u     *service.Service
	authn *auth.Auth
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server     *bm.ServerConfig
			AuthClient *warden.ClientConfig
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

	// 引入auth中间件
	authn = auth.New(&auth.Config{AuthClient: hc.AuthClient})

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
		g.POST("/sign_up", authn.AllowAny, SignUp)
		g.POST("/sign_in", authn.AllowAny, SignIn)
		g.GET("/info", authn.IsAuthenticated, Info)
		g.GET("/profile", authn.IsAuthenticated, Profile)
		g.POST("/password/update", authn.IsAuthenticated, SetPassword)
		g.POST("/sex/update", authn.IsAuthenticated, SetSex)
		g.POST("/face/update", authn.IsAuthenticated, SetFace)
		g.POST("/email/update", authn.IsAuthenticated, SetEmail)
		g.POST("/phone/update", authn.IsAuthenticated, SetPhone)

	}
}

func ping(ctx *bm.Context) {
	if err := u.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
