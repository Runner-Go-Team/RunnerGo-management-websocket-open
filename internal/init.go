package internal

import (
	"RunnerGo-management/internal/pkg/biz/log"
	"RunnerGo-management/internal/pkg/biz/proof"
	"RunnerGo-management/internal/pkg/conf"
	"RunnerGo-management/internal/pkg/dal"
	"go.uber.org/zap"
)

func InitProjects(readConfMode int, configFile string) {
	if readConfMode == 1 {
		conf.MustInitConfByEnv()
	} else {
		conf.MustInitConf(configFile)
	}
	dal.MustInitMySQL()
	dal.MustInitMongo()
	//dal.MustInitElasticSearch()
	proof.MustInitProof()
	//dal.MustInitGRPC()
	dal.MustInitRedis()
	dal.MustInitRedisForReport()
	dal.MustInitBigCache()
	// 初始化logger
	zap.S().Debug("初始化logger")
	log.InitLogger()

	//// 初始化redis客户端
	//if err := dal.InitRedisClient(
	//	conf.Conf.Redis.Address,
	//	conf.Conf.Redis.Password,
	//	int64(conf.Conf.Redis.DB),
	//); err != nil {
	//	panic("redis 连接失败")
	//}
}
