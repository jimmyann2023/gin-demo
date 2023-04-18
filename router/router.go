package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type RouteType = func(PublicRouter *gin.RouterGroup, AuthRouter *gin.RouterGroup)

var (
	RouteArray []RouteType
)

func RegisterRouter(fn RouteType) {
	if fn == nil {
		return
	}
	RouteArray = append(RouteArray, fn)
}

func InitRouter() {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	r := gin.Default()

	PublicRouter := r.Group("/api/v1/public")
	AuthRouter := r.Group("/api/v1/")
	InitBasePlatformRoutes()

	for _, fnRegisterRoute := range RouteArray {
		fnRegisterRoute(PublicRouter, AuthRouter)
	}

	Port := viper.GetString("server.port")
	if Port == "" {
		Port = "8999"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// TODO: 记录日志
			fmt.Println(fmt.Sprintf("Start Server Error: %s", err.Error()))
		}
		//fmt.Println(fmt.Sprintf("Srart Server Listen: %s", Port))
	}()
	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctx); err != nil {
		// TODO:记录日志
		fmt.Printf("Stop Server Error: %s\n", err.Error())
		return
	}
	fmt.Println("Stop Server Success")
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}
