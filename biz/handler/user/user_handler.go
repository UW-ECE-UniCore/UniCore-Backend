package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	UserService "unicore/biz/service/user"
	"unicore/pkg/errno"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var req UserRegisterRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(
			consts.StatusOK,
			UserRegisterResponse{
				StatusCode: errno.WrongFormat,
				StatusMSG:  err.Error(),
			},
		)
	} else {
		_, err = UserService.UserRegister(req.Email, req.UserName, req.Password)
		if err != nil {
			c.JSON(
				consts.StatusOK,
				UserRegisterResponse{
					StatusCode: errno.WrongFormat,
					StatusMSG:  err.Error(),
				},
			)
		} else {
			c.JSON(
				consts.StatusOK,
				UserRegisterResponse{
					StatusCode: errno.SuccessCode,
					StatusMSG:  errno.SuccessMsg,
				},
			)
		}
	}
}

func UserLogin(ctx context.Context, c *app.RequestContext) {
	email := c.Query("email")
	password := c.Query("password")
	user, err := UserService.UserLogin(email, password)
	if err != nil {
		errResp := err.(errno.CustomResponse)
		c.JSON(
			consts.StatusOK,
			UserLoginResponse{
				StatusCode: errResp.StatusCode,
				StatusMSG:  errResp.Message,
			},
		)
	} else {
		c.JSON(
			consts.StatusOK,
			UserLoginResponse{
				StatusCode: errno.SuccessCode,
				StatusMSG:  errno.SuccessMsg,
				UserName:   user.UserName,
				Email:      user.Email,
				UserID:     user.ID,
				Avatar:     user.Avatar,
				Signature:  user.Signature,
				School:     user.School,
			},
		)
	}

}
