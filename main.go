package main

import (
	"github.com/jimmyann2023/Gin/cmd"
)

func main() {
	defer cmd.Clean()
	cmd.Start()
}
