package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"unicore/biz/handler/user"
)

func Register(r *server.Hertz) {
	_user := r.Group("/user")
	{
		{
			_login := _user.Group("/login")
			_login.GET("/", user.UserLogin)
		}
		{
			_register := _user.Group("/register")
			_register.POST("/", user.UserRegister)
		}
	}
}
