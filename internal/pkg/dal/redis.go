package dal

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/conf"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func MustInitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Conf.Redis.Address,
		Password: conf.Conf.Redis.Password,
		DB:       conf.Conf.Redis.DB,
	})
}

func GetRDB() *redis.Client {
	return rdb
}
