package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/ecode"
)

func (s *Service) Token(ctx context.Context, token string) (mid int64, err error) {
	if token == "" {
		return
	}

	if mid, err = s.dao.GetMidByToken(ctx, token); err != nil {
		return
	}

	if mid == 0 {
		err = ecode.AccessDenied
	}
	return
}

func (s *Service) SetToken(ctx context.Context, token string, mid int64) (err error) {
	if token == "" {
		return
	}

	if err = s.dao.SetToken(ctx, token, mid); err != nil {
		return
	}

	return
}

func (s *Service) GetToken(ctx context.Context, mid int64) (token string, err error) {
	if mid == 0 {
		return
	}

	if token, err = s.dao.GetTokenByMid(mid); err != nil {
		return
	}

	return
}
