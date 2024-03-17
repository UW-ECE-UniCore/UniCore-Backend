package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"unicore/biz/handler/user"
	"unicore/biz/mw"
)

func Register(r *server.Hertz) {
	_user := r.Group("/user")
	{
		{
			_login := _user.Group("/login")
			_login.POST("/", mw.JwtMiddleware.LoginHandler)
		}
		{
			_register := _user.Group("/register")
			_register.POST("/", user.UserRegister)
		}
		{
			_getInfo := _user.Group("/get-info")
			_getInfo.GET("/", user.GetUserInfo)
		}
	}
}
