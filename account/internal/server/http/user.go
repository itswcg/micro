package http

import (
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/itswcg/micro/account/internal/model"
	"strconv"
)

func Info(ctx *bm.Context) {
	var (
		err    error
		mid    int64
		params = ctx.Request.Form
		midStr = params.Get("mid")
	)

	if mid, err = strconv.ParseInt(midStr, 10, 64); err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(u.GetInfo(ctx, mid))
}

func Profile(ctx *bm.Context) {
	var (
		err    error
		mid    int64
		params = ctx.Request.Form
		midStr = params.Get("mid")
	)

	if mid, err = strconv.ParseInt(midStr, 10, 64); err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(u.GetProfile(ctx, mid))
}

func SignUp(ctx *bm.Context) {
	type Request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	var (
		req Request
		err error
	)

	if err = ctx.Bind(&req); err != nil {
		return
	}

	// check name
	pass := u.FilterName(ctx, req.Name)
	if pass == false {
		ctx.JSON(nil, ecode.RequestErr)
	}

	token, err := u.GenerateToken(ctx)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	mid, err := u.AddInfo(ctx, req.Name, req.Password)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	err = u.SetName(ctx, req.Name, mid)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	// set token to redis
	err = u.SetToken(ctx, token, mid)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	info, err := u.GetProfile(ctx, mid)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	res := model.TokenInfo{Info: model.Info{Mid: info.Mid, Name: info.Name, Sex: info.Sex, Face: info.Face}, Token: token}

	ctx.JSON(res, nil)
}

func SignIn(ctx *bm.Context) {
	type Request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	var (
		err error
		req Request
	)

	if err = ctx.Bind(&req); err != nil {
		return
	}

	// 通过name 获取mid
	mid, err := u.GetMidByName(ctx, req.Name)
	if err != nil || mid == 0 {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	pass, err := u.CheckPassword(ctx, mid, req.Password)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	if pass == false {
		ctx.JSON(nil, ecode.Unauthorized)
		return
	}

	// Todo delete old token
	token, err := u.GenerateToken(ctx)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	err = u.SetToken(ctx, token, mid)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	info, err := u.GetProfile(ctx, mid)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	res := model.TokenInfo{Info: model.Info{Mid: info.Mid, Name: info.Name, Sex: info.Sex, Face: info.Face}, Token: token}

	ctx.JSON(res, nil)
}

func SetEmail(ctx *bm.Context) {
	type Request struct {
		Email string `json:"email"`
	}

	var (
		err     error
		req     Request
		mid, ok = ctx.Get("mid")
	)

	if err = ctx.Bind(&req); err != nil {
		return
	}

	if !ok {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}
	// check email

	err = u.SetEmail(ctx, mid.(int64), req.Email)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(nil, ecode.OK)
}

func SetPhone(ctx *bm.Context) {
	type Request struct {
		Phone string `json:"phone"`
	}

	var (
		err     error
		req     Request
		mid, ok = ctx.Get("mid")
	)

	if err = ctx.Bind(&req); err != nil {
		return
	}

	if !ok {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}
	// check phone

	err = u.SetPhone(ctx, mid.(int64), req.Phone)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(nil, ecode.OK)
}

func SetPassword(ctx *bm.Context) {
	type Request struct {
		Password string `json:"password"`
	}

	var (
		err     error
		req     Request
		mid, ok = ctx.Get("mid")
	)

	if err = ctx.Bind(&req); err != nil {
		return
	}

	if !ok {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	err = u.SetPassword(ctx, mid.(int64), req.Password)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(nil, ecode.OK)
}

func SetSex(ctx *bm.Context) {
	type Request struct {
		Sex string `json:"sex"`
	}

	var (
		err     error
		req     Request
		mid, ok = ctx.Get("mid")
	)

	if err = ctx.Bind(&req); err != nil {
		return
	}

	if !ok {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	err = u.SetSex(ctx, mid.(int64), req.Sex)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(nil, ecode.OK)
}

func SetFace(ctx *bm.Context) {
	type Request struct {
		Face string `json:"face"`
	}

	var (
		err     error
		req     Request
		mid, ok = ctx.Get("mid")
	)

	if err = ctx.Bind(&req); err != nil {
		return
	}

	if !ok {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	err = u.SetFace(ctx, mid.(int64), req.Face)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(nil, ecode.OK)
}
