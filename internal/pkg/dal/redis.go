package dal

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/conf"
	"github.com/go-redis/redis/v8"
	"strings"
)

var rdb *redis.ClusterClient

func MustInitRedis() {
	rdb = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split(conf.Conf.Redis.ClusterAddress, ";"),
		Password: conf.Conf.Redis.Password,
	})
}

func GetRDB() *redis.ClusterClient {
	return rdb
}
