package dal

import (
	"github.com/jichenssg/tikmall/app/auth/biz/dal/mysql"
	"github.com/jichenssg/tikmall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
