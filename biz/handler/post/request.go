package post

import "unicore/biz/dal/mysql"

type CreatePostRequest struct {
	Content   string `json:"content"`
	Title     string `json:"title"`
	CreatorID string `json:"user_id"`
	Type      string `json:"type"`
}

type CreatePostResponse struct {
	StatusCode int32      `json:"status_code"`
	Message    string     `json:"message"`
	Post       mysql.Post `json:"post"`
}

type GetPostResponse struct {
	StatusCode int32        `json:"status_code"`
	Message    string       `json:"message"`
	PostList   []mysql.Post `json:"post_list"`
}
