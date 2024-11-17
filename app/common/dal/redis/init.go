package redis

import (
	"fmt"

	"github.com/jichenssg/tikmall/app/common/obs"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func Init(host string, port int, db int, password string) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d",host, port),
		Password: password,
		DB:       db,
	})

	if err := redisotel.InstrumentTracing(rdb, redisotel.WithTracerProvider(obs.GetTraceProvider())); err != nil {
		panic(err)
	}
}

func GetClient() *redis.Client {
	return rdb
}
