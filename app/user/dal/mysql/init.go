package mysql

import (
	"fmt"
	"os"

	"github.com/jichenssg/tikmall/app/common/obs"
	"github.com/jichenssg/tikmall/app/user/config"
	"github.com/jichenssg/tikmall/app/user/dal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}

func Init() {
	var err error
	conf := config.GetConf().Mysql
	dsn := fmt.Sprintf(conf.Dsn, conf.User, conf.Password, conf.Host, conf.Port, conf.Database)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	if os.Getenv("GO_ENV") != "prod" {
		db.AutoMigrate(&model.User{})
	}

	// use otel tracing
	if err := db.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(obs.GetTraceProvider()))); err != nil {
		panic(err)
	}
}
