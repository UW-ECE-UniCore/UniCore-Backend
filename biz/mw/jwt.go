package mw

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
	"unicore/biz/dal/mysql"
	utils2 "unicore/pkg/utils"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "email"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"message": "success",
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}
			if err := c.Bind(&loginStruct); err != nil {
				return nil, err
			}
			users, err := mysql.CheckUser(loginStruct.Email, utils2.HashPassword(loginStruct.Password))
			if err != nil {
				return nil, err
			}
			if len(*users) == 0 {
				return nil, errors.New("user already exists or wrong password")
			}

			return &mysql.User{
				Email:    (*users)[0].Email,
				Password: (*users)[0].Password,
			}, nil
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &mysql.User{
				Email: claims[IdentityKey].(string),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*mysql.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
