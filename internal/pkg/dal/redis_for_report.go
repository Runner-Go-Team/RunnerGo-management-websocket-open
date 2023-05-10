package dal

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/conf"
	"github.com/go-redis/redis/v8"
)

var rdbReport *redis.Client

func MustInitRedisForReport() {
	rdbReport = redis.NewClient(&redis.Options{
		Addr:     conf.Conf.RedisReport.Address,
		Password: conf.Conf.RedisReport.Password,
		DB:       conf.Conf.RedisReport.DB,
	})
}

func GetRDBForReport() *redis.Client {
	return rdbReport
}
