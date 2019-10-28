package grpc

import (
	"context"
	"github.com/itswcg/micro/account/api"
)

func (s *UserServer) Info(ctx context.Context, req *api.MidReq) (res *api.InfoReply, err error) {
	info, err := s.u.GetInfo(ctx, req.Mid)
	if err != nil {
		return nil, err
	}
	return &api.InfoReply{Info: info}, nil
}

func (s *UserServer) Profile(ctx context.Context, req *api.MidReq) (res *api.ProfileReply, err error) {
	profile, err := s.u.GetProfile(ctx, req.Mid)
	if err != nil {
		return nil, err
	}
	return &api.ProfileReply{Profile: profile}, nil
}

func (s *UserServer) Token(ctx context.Context, req *api.TokenReq) (res *api.TokenReply, err error) {
	mid, err := s.u.Token(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	return &api.TokenReply{Mid: mid}, nil
}
