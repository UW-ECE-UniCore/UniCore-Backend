package dal

import (
	"unicore/biz/dal/mysql"
	"unicore/biz/mw/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
