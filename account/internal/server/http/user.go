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
		err error
		mid int64
	)

	params := ctx.Request.Form
	midStr := params.Get("mid")

	if mid, err = strconv.ParseInt(midStr, 10, 64); err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(u.GetProfile(ctx, mid))
}

func SignUp(ctx *bm.Context) {
	var (
		err      error
		params   = ctx.Request.Form
		name     = params.Get("name")
		password = params.Get("password")
	)

	// check name ? 判断唯一,大数据?
	token, err := u.GenerateToken(ctx)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	mid, err := u.AddInfo(ctx, name, password)
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
	var (
		err    error
		params = ctx.Request.Form
		//name     = params.Get("name")
		password = params.Get("password")
	)

	// 通过name 获取mid

	pass, err := u.CheckPassword(ctx, 1000, password)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	if pass == false {
		ctx.JSON(nil, ecode.Unauthorized)
		return
	}

	ctx.JSON(nil, err)
}

func SetEmail(ctx *bm.Context) {
	var (
		err     error
		params  = ctx.Request.Form
		mid, ok = ctx.Get("mid")
		email   = params.Get("email")
	)

	if !ok {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}
	// check email

	err = u.SetEmail(ctx, mid.(int64), email)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(nil, ecode.OK)
}

func SetPhone(ctx *bm.Context) {
	var (
		err     error
		params  = ctx.Request.Form
		mid, ok = ctx.Get("mid")
		phone   = params.Get("phone")
	)

	if !ok {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}
	// check phone

	err = u.SetPhone(ctx, mid.(int64), phone)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(nil, ecode.OK)
}

func SetPassword(ctx *bm.Context) {
	var (
		err      error
		params   = ctx.Request.Form
		mid, ok  = ctx.Get("mid")
		password = params.Get("password")
	)

	if !ok {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	err = u.SetPassword(ctx, mid.(int64), password)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(nil, ecode.OK)
}

func SetSex(ctx *bm.Context) {
	var (
		err     error
		params  = ctx.Request.Form
		mid, ok = ctx.Get("mid")
		sex     = params.Get("sex")
	)

	if !ok {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	err = u.SetSex(ctx, mid.(int64), sex)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(nil, ecode.OK)
}

func SetFace(ctx *bm.Context) {
	var (
		err     error
		params  = ctx.Request.Form
		mid, ok = ctx.Get("mid")
		face    = params.Get("face")
	)

	if !ok {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	err = u.SetFace(ctx, mid.(int64), face)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(nil, ecode.OK)
}
