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
	"net/http"
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

	// 创建监听 CTR + C , 应用退出信号的上下文
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

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

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: r,
	}

	go func() {
		global.Logger.Info(fmt.Sprintf("Start Listen: %s", Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Start Server Error: %s", err.Error()))
			return
		}
	}()
	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("Stop Server Error: %s", err.Error()))
		return
	}
	global.Logger.Info("Stop Server Success")
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}
