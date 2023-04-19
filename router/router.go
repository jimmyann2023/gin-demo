package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jimmyann2023/Gin/docs"
	"github.com/jimmyann2023/Gin/global"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type RouteType = func(PublicRouter *gin.RouterGroup, AuthRouter *gin.RouterGroup)

var (
	RouteArray []RouteType
)

// RegisterRouter 注册路由回调函数
func RegisterRouter(fn RouteType) {
	if fn == nil {
		return
	}
	RouteArray = append(RouteArray, fn)
}

// InitRouter 初始化系统路由
func InitRouter() {

	// 初始化 gin 框架,并注册相关路由
	r := gin.Default()

	PublicRouter := r.Group("/api/v1/public")
	AuthRouter := r.Group("/api/v1/")
	InitBasePlatformRoutes()

	for _, fnRegisterRoute := range RouteArray {
		fnRegisterRoute(PublicRouter, AuthRouter)
	}

	// 集成 swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	Port := viper.GetString("server.port")
	if Port == "" {
		Port = "8999"
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: r,
	}

	// 启动一个 goroutine 来开启web服务，避免主线程的信号监听被阻塞
	go func() {
		global.Logger.Info(fmt.Sprintf("Start Listen: %s", Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Start Sever Error: %s", err.Error()))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("Stop Server Error: %s", err.Error()))
	}
	global.Logger.Info("Sever exiting")

}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}
