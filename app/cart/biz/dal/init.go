package dal

import (
	"github.com/jichenssg/tikmall/app/cart/biz/dal/mysql"
	"github.com/jichenssg/tikmall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
