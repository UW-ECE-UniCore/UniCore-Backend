package mysql

import (
	"unicore/pkg/constants"
)

type User struct {
	ID              int64  `json:"id"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
}

func (User) TableName() string {
	return constants.UserTableName
}

func CreateUser(user *User) (int64, error) {
	err := DB.Create(user).Error
	if err != nil {
		return -1, err
	}
	return user.ID, nil
}

func GetAllUsers() (*[]User, error) {
	var UserList []User
	err := DB.Find(&UserList).Error
	if err != nil {
		return nil, err
	}
	return &UserList, nil
}
