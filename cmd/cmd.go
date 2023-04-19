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
	conf.InitConfig()
	global.Logger = conf.InitLogger()
	router.InitRouter()

	db, err := conf.InitDB()

	global.DB = db

	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
}

func Clean() {
	fmt.Println("====== clean ======")
}
