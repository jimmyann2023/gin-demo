package cmd

import (
	"fmt"
	"github.com/jimmyann2023/Gin/conf"
	"github.com/jimmyann2023/Gin/global"
	"github.com/jimmyann2023/Gin/router"
)

func Start() {
	conf.InitConfig()
	global.Logger = conf.InitLogger()
	router.InitRouter()
}

func Clean() {
	fmt.Println("====== clean ======")
}
