package errno

import "fmt"

const (
	SuccessCode         = 0
	WrongFormat         = 1
	MySQLQueryErr       = 2
	NotRegisteredErr    = 3
	WrongPasswordErr    = 4
	MySQLInsertErr      = 5
	LikeAlreadyExistErr = 6
)

const (
	SuccessMsg = "Success"
	//WrongFormatMSG     = "Input data can not be binded"
	MysqlQueryErrMSG      = "Query through mysql failed"
	EmailRegisteredErrMSG = "User's email has been registered"
	NotRegisteredErrMSG   = "This email has not been registered"
	WrongPasswordErrMSG   = "Wrong password input"

	LikeAlreadyExistErrMSG = "like already exists"
	//ServerErrMsg             = "Service is unable to start successfully"
	//ParamErrMsg              = "Wrong Parameter has been given"
	//UserIsNotExistErrMsg     = "user is not exist"
	//PasswordIsNotVerifiedMsg = "username or password not verified"
	//FavoriteActionErrMsg     = "favorite add failed"
	//
	//MessageAddFailedErrMsg    = "message add failed"
	//FriendListNoPermissionMsg = "You can't query his friend list"
	//VideoIsNotExistErrMsg     = "video is not exist"
	//CommentIsNotExistErrMsg   = "comment is not exist"
)

type CustomResponse struct {
	Message    string
	StatusCode int32
}

func (r CustomResponse) Error() string {
	return fmt.Sprintf("Error: %s,", r.Message)
}
