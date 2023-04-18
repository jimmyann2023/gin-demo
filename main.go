package main

import (
	"github.com/jimmyann2023/Gin/cmd"
)

//@title gin-sagger-demo 开发测试
//@version 1.0
//@description gin 框架的开发demo

func main() {
	defer cmd.Clean()
	cmd.Start()
}
