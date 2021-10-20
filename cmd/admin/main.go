package main

import (
	"guildchat-oss/internal"
	"guildchat-oss/internal/router"

	"github.com/gin-gonic/gin"
)

var (
	release bool = true
)

func main() {

	//判断是否编译线上版本
	if release {
		gin.SetMode(gin.ReleaseMode)
	}

	app := internal.Application{
		Route: router.Init(),
	}
	app.Run()

}