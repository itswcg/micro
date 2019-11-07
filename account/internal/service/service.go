package service

import (
	"context"
	"errors"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/itswcg/micro/account/internal/dao"
	authrpc "github.com/itswcg/micro/auth/api"
	leafrpc "github.com/itswcg/micro/leaf/api"
	"github.com/itswcg/micro/middleware/snowflake"
	"github.com/prometheus/common/log"
	"strconv"
)

// Service service.
type Service struct {
	ac    *paladin.Map
	dao   dao.Dao
	lcSrv leafrpc.LeafClient
	acSrv authrpc.AuthClient
}

// New new a service and return.
func New() (s *Service) {
	type Gc struct {
		LeafClient *warden.ClientConfig
		AuthClient *warden.ClientConfig
	}
	var (
		ac = new(paladin.TOML)
		gc = new(Gc) // new 返回指针, slice,map,chan使用make
	)

	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}

	if err := paladin.Get("grpc.toml").UnmarshalTOML(gc); err != nil {
		panic(err)
	}

	lcSrv, err := leafrpc.NewClient(gc.LeafClient)
	if err != nil {
		panic(err)
	}

	acSrv, err := authrpc.NewClient(gc.AuthClient)
	if err != nil {
		panic(err)
	}

	s = &Service{
		ac:    ac,
		dao:   dao.New(),
		lcSrv: lcSrv,
		acSrv: acSrv,
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
	LeafReply, err := s.lcSrv.NextID(ctx, LeafReq)
	if err != nil {
		log.Error("s.lc NextID(%v) err(%v)", LeafReq, err)
		return
	}
	return LeafReply.Id, nil
}

// Hash password
func (s *Service) GeneratePassword(ctx context.Context, password string) (hash_password string, err error) {
	hash_password = password + "test"
	return
}

// Check password
func (s *Service) CheckPassword(ctx context.Context, mid int64, password string) (pass bool, err error) {
	hash_password, err := s.GeneratePassword(ctx, password)
	if err != nil {
		return
	}
	db_password, err := s.dao.GetPassword(ctx, mid)
	if err != nil {
		return
	}
	if hash_password == db_password {
		pass = true
	}
	return
}

// Generate unique token
func (s *Service) GenerateToken(ctx context.Context) (token string, err error) {
	// 实现
	workId, err := s.ac.Get("snowflake_worker_id").Int64()
	if err != nil {
		return
	}

	work, err := snowflake.New(workId)
	if err != nil {
		return
	}

	ID := work.GetID()
	return strconv.FormatInt(ID, 10), nil
}

// Set Token to auth
func (s *Service) SetToken(ctx context.Context, token string, mid int64) (err error) {
	SetTokenReq := &authrpc.SetTokenReq{
		Token: token,
		Mid:   mid,
	}
	SetTokenReply, err := s.acSrv.SetToken(ctx, SetTokenReq)
	if err != nil {
		return
	}
	if SetTokenReply.Success == true {
		return nil
	} else {
		return errors.New("SetToken error")
	}
}
