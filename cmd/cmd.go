package cmd

import (
	"fmt"
	"github.com/jimmyann2023/Gin/conf"
	"github.com/jimmyann2023/Gin/global"
	"github.com/jimmyann2023/Gin/router"
	"github.com/jimmyann2023/Gin/utils"
)

func Start() {
	var initErr error

	// 初始化 系统配置
	conf.InitConfig()

	// 初始化日志
	global.Logger = conf.InitLogger()

	// 初始化 mysql
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// 初始化 redis
	redisClient, err := conf.InitRedis()
	global.RedisClient = redisClient

	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// 初始化过程中，遇到错误的最终处理
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}

	// 初始化路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("====== clean ======")
}
