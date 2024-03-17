package user

import (
	"unicore/biz/dal/mysql"
	"unicore/pkg/errno"
	"unicore/pkg/utils"
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
			Message: errno.EmailRegisteredErrMSG,
		}
	}
	user = &mysql.User{
		Email:    email,
		UserName: username,
		Password: utils.HashPassword(password),
	}
	_, err = mysql.CreateUser(user)
	if err != nil {
		return nil, errno.CustomResponse{
			Message: err.Error(),
		}
	}
	return user, nil
}

//func UserLogin(email, password string) (*mysql.User, error) {
//	user, err := mysql.GetUserByEmail(email)
//	if err != nil {
//		return nil, errno.CustomResponse{
//			Message:    errno.MysqlQueryErrMSG,
//			StatusCode: errno.MySQLQueryErr,
//		}
//	}
//	if user == nil {
//		return nil, errno.CustomResponse{
//			Message:    errno.NotRegisteredErrMSG,
//			StatusCode: errno.NotRegisteredErr,
//		}
//	}
//	if user.Password == utils.HashPassword(password) {
//		return user, nil
//	} else {
//		return nil, errno.CustomResponse{
//			Message:    errno.WrongPasswordErrMSG,
//			StatusCode: errno.WrongPasswordErr,
//		}
//	}
//}
