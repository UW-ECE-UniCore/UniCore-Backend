package post

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	PostService "unicore/biz/service/post"
	"unicore/pkg/errno"
)

func CreatePost(ctx context.Context, c *app.RequestContext) {
	var req CreatePostRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(
			consts.StatusOK,
			CreatePostResponse{
				StatusCode: errno.WrongFormat,
				StatusMSG:  err.Error(),
			},
		)
	} else {
		userID, _ := strconv.Atoi(req.CreatorID)
		post, err := PostService.CreatePost(req.Content, int64(userID))
		if err != nil {
			c.JSON(
				consts.StatusOK,
				CreatePostResponse{
					StatusCode: errno.MySQLInsertErr,
					StatusMSG:  err.Error(),
				},
			)
		} else {
			c.JSON(
				consts.StatusOK,
				CreatePostResponse{
					StatusCode: errno.SuccessCode,
					StatusMSG:  errno.SuccessMsg,
					Post:       *post,
				},
			)
		}
	}
}

func GetPostListByUser(ctx context.Context, c *app.RequestContext) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(
			consts.StatusOK,
			GetPostResponse{
				StatusCode: errno.WrongFormat,
				StatusMSG:  err.Error(),
			},
		)
	}
	postList, err := PostService.GetPostByUser(int64(userID))
	if err != nil {
		c.JSON(
			consts.StatusOK,
			GetPostResponse{
				StatusCode: errno.MySQLQueryErr,
				StatusMSG:  err.Error(),
			},
		)
	} else {
		c.JSON(
			consts.StatusOK,
			GetPostResponse{
				StatusCode: errno.SuccessCode,
				StatusMSG:  errno.SuccessMsg,
				PostList:   *postList,
			},
		)
	}
}
