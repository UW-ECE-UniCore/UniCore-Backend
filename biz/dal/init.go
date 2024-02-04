package dal

import (
	"unicore/biz/dal/mysql"
)

func Init() {
	mysql.Init()
	//redis.Init()
}
