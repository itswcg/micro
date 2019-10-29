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
