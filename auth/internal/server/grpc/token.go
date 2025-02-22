package grpc

import (
	"context"
	"github.com/itswcg/micro/auth/api"
)

func (s *AuthServer) Token(ctx context.Context, req *api.TokenReq) (res *api.TokenReply, err error) {
	mid, err := s.t.Token(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	return &api.TokenReply{Mid: mid}, nil
}

func (s *AuthServer) SetToken(ctx context.Context, req *api.SetTokenReq) (res *api.SetTokenReply, err error) {
	err = s.t.SetToken(ctx, req.Token, req.Mid)
	if err != nil {
		return nil, err
	}

	return &api.SetTokenReply{Success: true}, nil
}

func (s *AuthServer) GetToken(ctx context.Context, req *api.GetTokenReq) (res *api.GetTokenReply, err error) {
	token, err := s.t.GetToken(ctx, req.Mid)
	if err != nil {
		return nil, err
	}

	return &api.GetTokenReply{Token: token}, nil
}
