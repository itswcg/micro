package grpc

import (
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	pb "github.com/itswcg/micro/auth/api"
	"github.com/itswcg/micro/auth/internal/service"
)

type AuthServer struct {
	t *service.Service
}

// New new a grpc server.
func New(svc *service.Service) *warden.Server {
	var rc struct {
		Server *warden.ServerConfig
	}
	if err := paladin.Get("grpc.toml").UnmarshalTOML(&rc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	ws := warden.NewServer(rc.Server)
	pb.RegisterAuthServer(ws.Server(), &AuthServer{t: svc})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
