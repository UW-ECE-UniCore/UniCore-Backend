package mysql

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCreateUser(t *testing.T) {
	Init()
	u := &User{
		UserName:  "zjh",
		Email:     "jhz.travis@outlook.com",
		Password:  "277879201aA",
		Signature: "孟夏正需雨，一洗北晨昏",
	}
	userID, err := CreateUser(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(strconv.FormatInt(userID, 10))
}

func TestGetAllUsers(t *testing.T) {
	Init()
	userList, _ := GetAllUsers()
	fmt.Println(userList)
}
