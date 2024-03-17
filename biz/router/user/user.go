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
			_register.POST("/verifyEmail", user.VerifyEmail)
			_register.POST("/", user.UserRegister)
		}
		{
			_getInfo := _user.Group("/getInfoByEmail")
			_getInfo.GET("/", user.GetUserInfoByEmail)
		}
	}
}
