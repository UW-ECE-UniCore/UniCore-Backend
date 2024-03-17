package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"unicore/biz/dal/mysql"
	UserService "unicore/biz/service/user"
	"unicore/pkg/errno"
	"unicore/pkg/utils"
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
		school := utils.GetSchool(req.Email)
		if school == -1 {
			c.JSON(
				consts.StatusOK,
				UserRegisterResponse{
					Message: "Can not register a user with invalid email address.",
				},
			)
		}
		_, err = UserService.UserRegister(req.Email, req.UserName, req.Password, school)
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

func GetUserInfoByEmail(ctx context.Context, c *app.RequestContext) {
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

func VerifyEmail(ctx context.Context, c *app.RequestContext) {
	var req VerifyEmailRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(
			consts.StatusOK,
			VerifyEmailResponse{
				Message: err.Error(),
			},
		)
	} else {
		school := utils.GetSchool(req.Email)
		if school == -1 {
			c.JSON(
				consts.StatusOK,
				VerifyEmailResponse{
					Message: "School not registered, you can wait or contact the administrator to add your school.",
				},
			)
		} else {
			c.JSON(
				consts.StatusOK,
				VerifyEmailResponse{
					School:  school,
					Message: errno.SuccessMsg,
				},
			)
		}
	}
}
