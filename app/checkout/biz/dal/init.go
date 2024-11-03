package dal

import (
	"github.com/jichenssg/tikmall/app/checkout/biz/dal/mysql"
	"github.com/jichenssg/tikmall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
