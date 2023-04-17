package cmd

import (
	"fmt"
	"github.com/jimmyann2023/Gin/conf"
	"github.com/jimmyann2023/Gin/router"
)

func Start() {
	conf.InitConfig()
	router.InitRouter()
}

func Clean() {
	fmt.Println("====== clean ======")
}
