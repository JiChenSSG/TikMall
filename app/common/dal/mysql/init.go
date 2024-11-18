package mysql

import (
	"fmt"

	"github.com/jichenssg/tikmall/app/common/obs"
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

func Init(dsn string, username string, password string, host string, port int, database string) {
	var err error
	dsn = fmt.Sprintf(dsn, username, password, host, port, database)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	// use otel tracing
	if err := db.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(obs.GetTraceProvider()))); err != nil {
		panic(err)
	}
}
