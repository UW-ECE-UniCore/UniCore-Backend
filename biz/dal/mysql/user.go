package mysql

import (
	"errors"
	"gorm.io/gorm"
	"unicore/pkg/constants"
)

type User struct {
	ID              int64  `json:"id"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	School          string `json:"school"`
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

func CheckUser(email, password string) (*[]User, error) {
	var UserList []User
	err := DB.Where("email = ? AND password = ?", email, password).Find(&UserList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &UserList, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() (*[]User, error) {
	var UserList []User
	err := DB.Find(&UserList).Error
	if err != nil {
		return nil, err
	}
	return &UserList, nil
}
