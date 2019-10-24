package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/itswcg/micro/account/internal/dao"
	leafrpc "github.com/itswcg/micro/leaf/api"
	"github.com/prometheus/common/log"
)

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
	lc  leafrpc.LeafClient
}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	var (
		gc struct {
			LeafClient *warden.ClientConfig
		}
	)

	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}

	if err := paladin.Get("grpc.toml").UnmarshalTOML(&gc); err != nil {
		panic(err)
	}

	lc, err := leafrpc.NewClient(gc.LeafClient)
	if err != nil {
		panic(err)
	}

	s = &Service{
		ac:  ac,
		dao: dao.New(),
		lc:  lc,
	}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}

// Rpc nextId
func (s *Service) NextID(ctx context.Context) (id int64, err error) {
	LeafReq := &leafrpc.LeafReq{
		BizTag: "account",
	}
	LeafReply, err := s.lc.NextID(ctx, LeafReq)
	if err != nil {
		log.Error("s.lc NextID(%v) err(%v)", LeafReq, err)
		return
	}
	return LeafReply.Id, nil
}
