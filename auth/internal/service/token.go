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
