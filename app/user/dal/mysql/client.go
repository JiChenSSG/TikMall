package mysql

import (
	"fmt"
	"os"
	"sync"

	"github.com/jichenssg/tikmall/app/user/config"
	"github.com/jichenssg/tikmall/app/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		db = Init()
	})
	return db
}

func Init() *gorm.DB {
	conf := config.GetConf().Mysql
	dsn := fmt.Sprintf(conf.Dsn, conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	if os.Getenv("GO_ENV") != "prod" {
		db.AutoMigrate(&model.User{})
	}

	return db
}
