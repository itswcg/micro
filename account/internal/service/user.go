package service

import (
	"account/api"
	"account/internal/model"
	"context"
	"github.com/bilibili/kratos/pkg/ecode"
)

func (s *Service) GetInfo(ctx context.Context, mid int64) (info *api.Info, err error) {
	// 缓存
	if mid < 1 {
		err = ecode.RequestErr
		return
	}
	if info, err := s.dao.GetInfo(ctx, mid); err != nil {
		return info, err
	}
	return
}

func (s *Service) GetProfile(ctx context.Context, mid int64) (profile *api.Profile, err error) {
	// 缓存
	if mid < 1 {
		err = ecode.RequestErr
		return
	}
	if profile, err := s.dao.GetProfile(ctx, mid); err != nil {
		return profile, err
	}
	return
}

func (s *Service) AddInfo(ctx context.Context, name, sex, face string) (err error) {
	var mid int64
	mid = 3
	info := &model.Info{
		Mid:  mid,
		Name: name,
		Sex:  sex,
		Face: face,
	}
	err = s.dao.AddInfo(ctx, info)
	return
}

func (s *Service) SetEmail(ctx context.Context) {

}

func (s *Service) SetPhone(ctx context.Context) {

}
