package post

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"unicore/biz/handler/post"
)

func Register(r *server.Hertz) {
	_post := r.Group("/post")
	{
		{
			_create := _post.Group("/create")
			_create.POST("/", append(mwFunc(), post.CreatePost)...)
		}
		{
			_getPost := _post.Group("/get-post")
			{
				_userPost := _getPost.Group("/user")
				_userPost.GET("/", append(mwFunc(), post.GetPostListByUser)...)
			}
			//{
			//	_schoolPost := _getPost.Group("/school")
			//	_schoolPost.GET("/", post.GetPostListBySchool)
			//}
			//{
			//	_publicPost := _getPost.Group("/public")
			//	_publicPost.GET("/", post.GetPublicPost)
			//}
		}
	}
}
