package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/itswcg/micro/account/api"
	"github.com/itswcg/micro/account/internal/model"
)

func (s *Service) GetInfo(ctx context.Context, mid int64) (info *api.Info, err error) {
	// 缓存
	if mid < 1 {
		err = ecode.RequestErr
		return
	}
	if info, err = s.dao.GetInfo(ctx, mid); err != nil {
		err = ecode.RequestErr
		return
	}
	return
}

func (s *Service) GetProfile(ctx context.Context, mid int64) (profile *api.Profile, err error) {
	// 缓存
	if mid < 1 {
		err = ecode.RequestErr
		return
	}
	if profile, err = s.dao.GetProfile(ctx, mid); err != nil {
		err = ecode.RequestErr
		return
	}
	return
}

func (s *Service) AddInfo(ctx context.Context, name, sex, face string) (mid int64, err error) {
	mid, err = s.NextID(ctx)
	if err != nil {
		return
	}

	info := &model.Info{
		Mid:  mid,
		Name: name,
		Sex:  sex,
		Face: face,
	}
	err = s.dao.AddInfo(ctx, info)
	return
}

func (s *Service) SetEmail(ctx context.Context, mid int64, email string) (err error) {
	err = s.dao.SetEmail(ctx, mid, email)
	return
}

func (s *Service) SetPhone(ctx context.Context, mid int64, phone string) (err error) {
	err = s.dao.SetPhone(ctx, mid, phone)
	return
}
