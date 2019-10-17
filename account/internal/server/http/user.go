package http

import (
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
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

func AddInfo(ctx *bm.Context) {
	var (
		err error
	)
	params := ctx.Request.Form
	name := params.Get("name")
	sex := params.Get("sex")
	face := params.Get("face")

	// check name
	// check face

	err = u.AddInfo(ctx, name, sex, face)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

}

func SetEmail(ctx *bm.Context) {

}

func SetPhone(ctx *bm.Context) {

}
