package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/itswcg/micro/account/api"
	"github.com/prometheus/common/log"
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

func (s *Service) AddInfo(ctx context.Context, name, password string) (mid int64, err error) {
	mid, err = s.NextID(ctx)
	if err != nil {
		return
	}

	hash_password, err := s.GeneratePassword(ctx, password)
	if err != nil {
		return
	}

	err = s.dao.AddInfo(ctx, mid, name, hash_password)
	return
}

func (s *Service) SetEmail(ctx context.Context, mid int64, email string) (err error) {
	err = s.dao.SetInfo(ctx, mid, "email", email)
	return
}

func (s *Service) SetPhone(ctx context.Context, mid int64, phone string) (err error) {
	err = s.dao.SetInfo(ctx, mid, "phone", phone)
	return
}

func (s *Service) SetSex(ctx context.Context, mid int64, sex string) (err error) {
	err = s.dao.SetInfo(ctx, mid, "sex", sex)
	return
}

func (s *Service) SetFace(ctx context.Context, mid int64, face string) (err error) {
	err = s.dao.SetInfo(ctx, mid, "face", face)
	return
}

func (s *Service) SetPassword(ctx context.Context, mid int64, password string) (err error) {
	hash_passowrd, err := s.GeneratePassword(ctx, password)
	if err != nil {
		return
	}

	err = s.dao.SetInfo(ctx, mid, "password", hash_passowrd)
	return
}

func (s *Service) FilterName(ctx context.Context, name string) (pass bool) {
	if name == "" {
		log.Error("name is null")
		pass = false
	}

	mid, err := s.GetMidByName(ctx, name)
	if err != nil {
		log.Error("getMidByName error(%v)", err)
		pass = false
	}

	if mid == 0 {
		pass = true
	}

	return
}

func (s *Service) SetName(ctx context.Context, name string, mid int64) (err error) {
	err = s.dao.SetName(ctx, name, mid)
	return
}

func (s *Service) GetMidByName(ctx context.Context, name string) (mid int64, err error) {
	mid, err = s.dao.GetMidByName(ctx, name)
	return
}
