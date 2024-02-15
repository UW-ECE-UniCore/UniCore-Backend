package mysql

import (
	"errors"
	"gorm.io/gorm"
	"unicore/pkg/constants"
)

type Post struct {
	PostID     int64  `json:post_id`
	CreatorID  int64  `json:creator_id`
	Content    string `json:content`
	Type       string `json:type`
	CreateTime string `json:create_time`
}

func (Post) TableName() string {
	return constants.PostTableName
}

func CreatePost(post *Post) (int64, error) {
	err := DB.Create(post).Error
	if err != nil {
		return -1, err
	}
	return post.PostID, nil
}

func GetPostByPostID(postID int64) (*Post, error) {
	var post Post
	err := DB.Where("post_id = ?", postID).First(&post).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &post, nil
}

func GetPostListByUserID(userID int64) (*[]Post, error) {
	var postList []Post
	res := DB.Order("create_time desc").Where("creator_id = ?", userID).Find(&postList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &postList, nil
		}
		return nil, res.Error
	}
	return &postList, nil
}
