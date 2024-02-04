package user

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	UserService "unicore/biz/service/user"
	"unicore/pkg/errno"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var req UserRegisterRequest
	err := c.Bind(&req)
	fmt.Println(req.Email)
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

//func UserLogin(ctx context.Context, c *app.RequestContext) {
//	var req userLoginRequest
//	err := c.BindAndValidate(&req)
//	if err != nil {
//
//	}
//
//}
