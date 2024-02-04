package user

import (
	"unicore/biz/dal/mysql"
	"unicore/pkg/errno"
)

func UserRegister(email string, username string, password string) (*mysql.User, error) {
	user, err := mysql.GetUserByEmail(email)
	if err != nil {
		return nil, errno.CustomResponse{
			Message: errno.MysqlQueryErrMSG,
		}
	}
	if user != nil {
		return nil, errno.CustomResponse{
			Message: errno.EmailRegisteredErr,
		}
	}
	user = &mysql.User{
		Email:    email,
		UserName: username,
		Password: password,
	}
	_, err = mysql.CreateUser(user)
	if err != nil {
		return nil, errno.CustomResponse{
			Message: err.Error(),
		}
	}
	return user, nil
}

//func UserLogin() (userID int64, err string) {
//
//}
