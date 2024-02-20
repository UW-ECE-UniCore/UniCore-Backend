package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"unicore/biz/router/like"
	"unicore/biz/router/post"
	"unicore/biz/router/user"
)

func RegisterRouters(r *server.Hertz) {
	user.Register(r)
	post.Register(r)
	like.Register(r)
}
