package router

import (
	"guildchat-oss/internal/middleware"
	"guildchat-oss/pkg/comment"
	"guildchat-oss/pkg/loggers/facade"
	"guildchat-oss/pkg/medium"
	"guildchat-oss/web"
	"path/filepath"
	"time"

	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	SwagHandler gin.HandlerFunc
)

func Init() *gin.Engine {

	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	prep(router)

	// router.Use(medium.GinLog(facade.NewZaplog("admin"), time.RFC3339, true), medium.RecoveryWithLog(facade.NewZaplog("admin"), true))
	router.Use(middleware.NotHttpStatusOk())
	router.Use(medium.GinLog(facade.NewRedisLog("admin"), time.RFC3339, true), medium.RecoveryWithLog(facade.NewRedisLog("admin"), true))
	/*****admin路由定义******/
	adminRouter := router.Group("/admin")

	AdminRouter(adminRouter)

	/***api路由定义****/
	apiRouter := router.Group("/api")

	ApiRouter(apiRouter)

	return router
}

func prep(router *gin.Engine) {
	var (
		rootPath   string
		uploadPath string
		err        error
	)

	rootPath, _ = comment.RootPath()

	uploadPath = rootPath + string(filepath.Separator) + "uploadfile"

	if SwagHandler != nil {
		router.GET("/swagger/*any", SwagHandler)
	}

	router.HTMLRender = web.LoadTemplates()

	router.StaticFS("/statics", web.StaticsFs)

	_, err = os.Stat(uploadPath)

	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(uploadPath, os.ModeDir)
			if err != nil {
				panic(err.Error())
			}
		}
	}

	router.StaticFS("/uploadfile", http.Dir(uploadPath))
}
