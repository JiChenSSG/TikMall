package dal

import (
	"github.com/jichenssg/tikmall/app/order/biz/dal/mysql"
	"github.com/jichenssg/tikmall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
