package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"unicore/biz/dal/mysql"
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
				Message:    err.Error(),
			},
		)
	} else {
		_, err = UserService.UserRegister(req.Email, req.UserName, req.Password)
		if err != nil {
			c.JSON(
				consts.StatusOK,
				UserRegisterResponse{
					StatusCode: errno.MySQLInsertErr,
					Message:    err.Error(),
				},
			)
		} else {
			c.JSON(
				consts.StatusOK,
				UserRegisterResponse{
					StatusCode: errno.SuccessCode,
					Message:    errno.SuccessMsg,
				},
			)
		}
	}
}

//func UserLogin(ctx context.Context, c *app.RequestContext) {
//	email := c.Query("email")
//	password := c.Query("password")
//	user, err := UserService.UserLogin(email, password)
//	if err != nil {
//		errResp := err.(errno.CustomResponse)
//		c.JSON(
//			consts.StatusOK,
//			UserLoginResponse{
//				StatusCode: errResp.StatusCode,
//				Message:    errResp.Message,
//			},
//		)
//	} else {
//		c.JSON(
//			consts.StatusOK,
//			UserLoginResponse{
//				StatusCode: errno.SuccessCode,
//				Message:    errno.SuccessMsg,
//				UserName:   user.UserName,
//				Email:      user.Email,
//				UserID:     user.ID,
//				Avatar:     user.Avatar,
//				Signature:  user.Signature,
//				School:     user.School,
//			},
//		)
//	}
//
//}

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	email := c.Query("email")
	user, err := mysql.GetUserByEmail(email)
	if err != nil {
		c.JSON(
			consts.StatusOK,
			GetUserInfoResponse{
				Message: err.Error(),
			},
		)
	} else if user == nil {
		c.JSON(
			consts.StatusOK,
			GetUserInfoResponse{
				Message: "No user found with email: " + email,
			},
		)
	} else {
		c.JSON(
			consts.StatusOK,
			GetUserInfoResponse{
				Message:   errno.SuccessMsg,
				UserName:  user.UserName,
				Email:     user.Email,
				UserID:    user.ID,
				Avatar:    user.Avatar,
				Signature: user.Signature,
				School:    user.School,
			},
		)
	}
}
