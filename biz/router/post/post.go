package post

import "github.com/cloudwego/hertz/pkg/app/server"

func Register(r *server.Hertz) {
	_post := r.Group("/post")
	{
		{
			_create := _post.Group("/create")
			_create.POST("/", post.CreatePost)
		}
		{
			_showPost := _post.Group("/showPost")
		}
	}
}
