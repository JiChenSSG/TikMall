package dal

import (
	"github.com/jichenssg/tikmall/app/product/biz/dal/mysql"
	"github.com/jichenssg/tikmall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
