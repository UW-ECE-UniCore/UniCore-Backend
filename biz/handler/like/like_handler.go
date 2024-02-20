package like

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	LikeService "unicore/biz/service/like"
	"unicore/pkg/errno"
)

func CheckLikeExistHandler(ctx context.Context, c *app.RequestContext) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	postID, err := strconv.Atoi(c.Query("post_id"))
	if err != nil {
		c.JSON(
			consts.StatusOK,
			CheckLikeExistResponse{
				StatusCode: errno.StrConvErr,
				Message:    err.Error(),
			},
		)
		return
	}
	res, err := LikeService.CheckLikeExistService(int64(userID), int64(postID))
	if err != nil {
		cusErr := err.(errno.CustomResponse)
		c.JSON(
			consts.StatusOK,
			CheckLikeExistResponse{
				StatusCode: cusErr.StatusCode,
				Message:    cusErr.Message,
			},
		)
		return
	}
	c.JSON(
		consts.StatusOK,
		CheckLikeExistResponse{
			StatusCode: errno.SuccessCode,
			Message:    errno.SuccessMsg,
			Existence:  res,
		},
	)
	return
}

func ActionHandler(ctx context.Context, c *app.RequestContext) {
	var req ActionRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(
			consts.StatusOK,
			ActionResponse{
				StatusCode: errno.WrongFormat,
				Message:    err.Error(),
			},
		)
		return
	}
	userID, err := strconv.Atoi(req.UserID)
	postID, err := strconv.Atoi(req.PostID)
	if err != nil {
		c.JSON(
			consts.StatusOK,
			ActionResponse{
				StatusCode: errno.StrConvErr,
				Message:    err.Error(),
			},
		)
		return
	}
	res, err := LikeService.ActionService(int64(userID), int64(postID))
	if err != nil {
		cusErr := err.(errno.CustomResponse)
		c.JSON(
			consts.StatusOK,
			ActionResponse{
				StatusCode: cusErr.StatusCode,
				Message:    cusErr.Message,
			},
		)
		return
	}
	c.JSON(
		consts.StatusOK,
		ActionResponse{
			StatusCode: errno.SuccessCode,
			Message:    errno.SuccessMsg,
			Existence:  res,
		},
	)
}
