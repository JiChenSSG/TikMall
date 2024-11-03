package dal

import (
	"github.com/jichenssg/tikmall/app/user/biz/dal/mysql"
	"github.com/jichenssg/tikmall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
