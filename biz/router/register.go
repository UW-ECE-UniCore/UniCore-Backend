package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"unicore/biz/router/user"
)

func RegisterRouters(r *server.Hertz) {
	user.Register(r)
}
