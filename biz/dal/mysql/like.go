package mysql

import "unicore/pkg/constants"

type Like struct {
	LikeID     int64  `json:"like_id"`
	FromUserID int64  `json:"from_user_id"`
	PostID     int64  `json:"post_id"`
	CreateTime string `json:"create_time"`
}

func (Like) TableName() string {
	return constants.LikeTableName
}

func AddLike(like *Like) (bool, error) {
	err := DB.Create(like).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteLike(like *Like) (bool, error) {
	err := DB.Where("post_id = ? AND from_user_id = ?", like.PostID, like.FromUserID).Delete(like).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func CheckLikeExist(userID, postID int64) (bool, error) {
	var sum int64
	err := DB.Model(&Like{}).Where("post_id = ? AND from_user_id = ?", postID, userID).Count(&sum).Error
	if err != nil {
		return false, err
	}
	if sum == 0 {
		return false, nil
	}
	return true, nil
}
