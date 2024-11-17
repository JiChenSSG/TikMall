package model

import (
	"context"
	"time"

	"github.com/jichenssg/tikmall/app/auth/config"
	"github.com/jichenssg/tikmall/app/auth/utils"
	"github.com/redis/go-redis/v9"
)

func CacheToken(rdb *redis.Client, ctx context.Context, token string, userID int64) error {
	ttl := config.GetConf().Token.RefreshExpire

	return rdb.Set(ctx, utils.TokenKey(userID), token, time.Duration(ttl)*time.Second).Err()
}

func ExtendToken(rdb redis.Client, ctx context.Context, userID int64) error {
	script := `
		if redis.call("EXISTS", KEYS[1]) == 1 then
			return redis.call("EXPIRE", KEYS[1], ARGV[1])
		else
			return 0
		end
	`

	ttl := config.GetConf().Token.RefreshExpire

	return rdb.Eval(ctx, script, []string{utils.TokenKey(userID)}, ttl).Err()
}

func DeleteToken(rdb *redis.Client, ctx context.Context, userID int64) error {
	return rdb.Del(ctx, utils.TokenKey(userID)).Err()
}
