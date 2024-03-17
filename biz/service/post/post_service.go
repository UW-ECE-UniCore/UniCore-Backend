package post

import (
	"unicore/biz/dal/mysql"
	"unicore/pkg/constants"
	"unicore/pkg/errno"
	"unicore/pkg/utils"
)

func CreatePost(content string, creatorID int64, school int16, title string) (*mysql.Post, error) {
	curTime := utils.GetCurrentTime()
	post := &mysql.Post{
		CreatorID:  creatorID,
		Title:      title,
		Type:       constants.TypePublic,
		Status:     constants.StatusCreate,
		School:     school,
		Content:    content,
		CreateTime: curTime,
	}
	_, err := mysql.CreatePost(post)
	if err != nil {
		return nil, errno.CustomResponse{
			Message: err.Error(),
		}
	}
	return post, nil
}

func GetPostByUser(userID int64) (*[]mysql.Post, error) {
	postList, err := mysql.GetPostListByUserID(userID)
	if err != nil {
		return nil, errno.CustomResponse{
			Message: errno.MysqlQueryErrMSG,
		}
	}
	return postList, nil
}

//func GetPostBySchool(userID int64) (*[]Post, error) {
//	postList, err := mysql.GetPostListBySchool(userID)
//	if err != nil {
//		return nil, errno.CustomResponse{
//			Message: errno.MysqlQueryErrMSG,
//		}
//	}
//	return postList, nil
//}
