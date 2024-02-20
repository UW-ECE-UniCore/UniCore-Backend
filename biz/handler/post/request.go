package post

import "unicore/biz/dal/mysql"

type CreatePostRequest struct {
	Content   string `json:"content"`
	CreatorID string `json:"user_id"`
}

type CreatePostResponse struct {
	StatusCode int32
	StatusMSG  string
	Post       mysql.Post
}

type GetPostResponse struct {
	StatusCode int32
	StatusMSG  string
	PostList   []mysql.Post
}
