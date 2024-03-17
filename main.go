package main

import (
	"context"
	"unicore/biz/mw"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"unicore/biz/dal"
)

func init() {
	dal.Init()
}

func main() {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8080"),
	)

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	mw.InitJwt()
	register(h)

	h.Spin()
}
