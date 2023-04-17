package cmd

import (
	"fmt"
	"github.com/jimmyann2023/Gin/conf"
)

func Start() {
	conf.InitConfig()
}

func Clean() {
	fmt.Println("====== clean ======")
}
