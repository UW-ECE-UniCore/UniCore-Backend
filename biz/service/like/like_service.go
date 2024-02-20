package like

import (
	"unicore/biz/dal/mysql"
	"unicore/pkg/errno"
)

//func AddNewLikeService(userID, postID int64) (bool, error) {
//	exist, err := mysql.CheckLikeExist(userID, postID)
//	if err != nil {
//		return false, errno.CustomResponse{
//			Message: errno.MysqlQueryErrMSG,
//		}
//	}
//	if exist {
//		return false, errno.CustomResponse{
//			Message: errno.LikeAlreadyExistErrMSG,
//		}
//	}
//	like := &mysql.Like{
//		FromUserID: userID,
//		PostID:     postID,
//	}
//	res, err := mysql.AddLike(like)
//	if err != nil {
//		return false, errno.CustomResponse{
//			Message: err.Error(),
//		}
//	}
//	return res, nil
//}

//func DeleteLike(userID, postID int64) (bool, error) {
//	like := &mysql.Like{
//		PostID:     postID,
//		FromUserID: userID,
//	}
//	res, err := mysql.DeleteLike(like)
//}

func CheckLikeExistService(userID, postID int64) (bool, error) {
	res, err := mysql.CheckLikeExist(userID, postID)
	if err != nil {
		return false, errno.CustomResponse{
			Message: err.Error(),
		}
	}
	return res, nil
}

func LikeActionService(userID, postID int64) (bool, error) {
	res, err := mysql.CheckLikeExist(userID, postID)
	if err != nil {
		return false, errno.CustomResponse{
			Message: err.Error(),
		}
	}
	if res {
		like := &mysql.Like{
			FromUserID: userID,
			PostID:     postID,
		}
		_, err := mysql.DeleteLike(like)
		if err != nil {
			return false, errno.CustomResponse{
				Message: err.Error(),
			}
		}
	} else {
		like := &mysql.Like{
			FromUserID: userID,
			PostID:     postID,
		}
		_, err := mysql.AddLike(like)
		if err != nil {
			return false, errno.CustomResponse{
				Message: err.Error(),
			}
		}
	}
	return true, nil
}
