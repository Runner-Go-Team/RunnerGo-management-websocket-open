package internal

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/log"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/proof"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/conf"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
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
	dal.MustInitBigCache()
	// 初始化logger
	zap.S().Debug("初始化logger")
	log.InitLogger()
}
