package main

import (
	"context"
	"fastgin/config" // 这个包会在运行 swag init 后自动生成
	"fastgin/internal/middleware"
	"fastgin/internal/repository"
	"fastgin/internal/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// @title FastGin API
// @version 1.0
// @description FastGin 服务API文档
// @host localhost:8080
// @BasePath /api
func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("配置初始化失败: %v", err)
	}

	// 初始化日志
	middleware.InitLogger()

	// 启动配置热重载
	config.WatchConfig()

	// 初始化数据库连接
	err, db := repository.InitDB()
	if err != nil {
		panic(err)
	}
	// 服务启动端口
	conf := config.GlobalConfig
	port := ":" + conf.Server.Port

	// 初始化 路由
	r := router.InitRouter(db, conf)

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	// 在独立的 goroutine 中启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			middleware.Logger.Fatal("监听失败", zap.Error(err))
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	middleware.Logger.Info("正在关闭服务器...")

	// 设置关闭超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		middleware.Logger.Fatal("服务器强制关闭", zap.Error(err))
	}

	// 等待数据库连接关闭
	sqlDB, err := db.DB()
	if err != nil {
		middleware.Logger.Error("获取数据库连接失败", zap.Error(err))
	} else {
		if err := sqlDB.Close(); err != nil {
			middleware.Logger.Error("关闭数据库连接失败", zap.Error(err))
		}
	}

	middleware.Logger.Info("服务器已优雅关闭")
}
