package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

	err := r.Run(fmt.Sprintf(":%s", Port))
	if err != nil {
		panic(fmt.Sprintf("Strart Server Error: %s", err.Error()))
	}
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}
