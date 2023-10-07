package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-omnibus/proof"

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/handler"
)

func RegisterRouter(r *gin.Engine) {
	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Upgrade", "Origin", "Connection", "Accept-Encoding", "Accept-Language", "Host", "x-requested-with", "CurrentTeamID"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(ginZap.Ginzap(proof.Logger.Z, time.RFC3339, true))

	r.Use(ginZap.RecoveryWithZap(proof.Logger.Z, true))

	// 探活接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// websocket相关接口
	// 独立报告页面接口
	websocket := r.Group("websocket")
	websocket.GET("index", handler.WebSocket)
}
