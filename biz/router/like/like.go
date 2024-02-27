package like

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	LikeHandler "unicore/biz/handler/like"
)

func Register(r *server.Hertz) {
	_like := r.Group("/like")
	{
		_check := _like.Group("/check")
		_check.GET("/", LikeHandler.CheckLikeExistHandler)
	}
	{
		_action := _like.Group("/action")
		_action.POST("/", LikeHandler.ActionHandler)
	}
}
