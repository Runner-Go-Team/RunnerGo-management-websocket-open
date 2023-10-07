package main

import (
	"flag"
	"fmt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/app/router"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/conf"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/handler"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var readConfMode int
var configFile string

func main() {
	flag.IntVar(&readConfMode, "m", 0, "读取环境变量还是读取配置文件")
	flag.StringVar(&configFile, "c", "./configs/dev.yaml", "app config file.")
	flag.Parse()

	internal.InitProjects(readConfMode, configFile)

	// 定期删除失效的websocket连接
	go func() {
		handler.CloseInvalidWbLink()
	}()

	// 主动给用户发送当前在运行中计划数量
	go func() {
		handler.PushRunningPlanCount()
	}()

	// 消费 UI 自动化结果
	go func() {
		handler.ConsumerUIEngineResult()
	}()

	// 创建 Gin 引擎实例
	engine := gin.Default()
	router.RegisterRouter(engine)

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Conf.Http.Port),
		Handler: engine,
	}

	// 启动 HTTP 服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务启动失败: %v", err)
		}
	}()

	// 优雅退出
	gracefulExit(server)
}

// 释放连接池资源，优雅退出
func gracefulExit(server *http.Server) {
	// 等待中断信号，然后优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("服务关闭中...")

	// 创建一个 5 秒的上下文，用于等待请求处理完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭服务器，等待现有的请求处理完成
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("服务关闭失败: %v", err)
	}

	log.Println("服务已经优雅的关闭~~~")
}
