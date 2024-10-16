package initialize

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PandaPy/pginer/template/initialize/config"
	"github.com/PandaPy/pginer/template/initialize/db"
	"github.com/PandaPy/pginer/template/initialize/logger"
	"github.com/PandaPy/pginer/template/initialize/validator"
	"github.com/PandaPy/pginer/template/router"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func RunWindowsServer() {

	logger.Init()    // 初始化日志
	config.Init()    // 初始化配置
	db.Init()        // 初始化数据库
	validator.Init() // 初始化验证器

	gin.SetMode(config.AppConfig.Server.Mode)

	r := gin.New()

	router.SetupRoutes(r)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Server.Listen),
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	color.Yellow("Server successfully running on port: %d", config.AppConfig.Server.Listen)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := server.Close(); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}
}
