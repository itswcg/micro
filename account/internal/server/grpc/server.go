package grpc

import (
	pb "account/api"
	"account/internal/service"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

type UserServer struct {
	u *service.Service
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
	pb.RegisterAccountServer(ws.Server(), &UserServer{u: svc})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
