package dal

import (
	"github.com/jichenssg/tikmall/app/payment/biz/dal/mysql"
	"github.com/jichenssg/tikmall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
