package main

import (
	"flag"
	"fmt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/app/router"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/conf"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/handler"
	"github.com/gin-gonic/gin"
)

var readConfMode int
var configFile string

func main() {
	flag.IntVar(&readConfMode, "m", 0, "读取环境变量还是读取配置文件")
	flag.StringVar(&configFile, "c", "./configs/dev.yaml", "app config file.")
	flag.Parse()

	internal.InitProjects(readConfMode, configFile)

	r := gin.New()
	router.RegisterRouter(r)
	// 定期删除失效的websocket连接
	go func() {
		handler.CloseInvalidWbLink()
	}()

	// 主动给用户发送当前在运行中计划数量
	go func() {
		handler.PushRunningPlanCount()
	}()

	if err := r.Run(fmt.Sprintf(":%d", conf.Conf.Http.Port)); err != nil {
		panic(err)
	}
}
