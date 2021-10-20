package router

import (

	"guildchat-oss/internal/controllers/admin"
	"guildchat-oss/internal/controllers/admin/setting"
	"guildchat-oss/internal/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AdminRouter(adminRouter *gin.RouterGroup) {

	//设置后台用户权限中间件
	store := cookie.NewStore([]byte("secret11111"))
	adminRouter.Use(sessions.Sessions("mysession", store))
	{
		adminRouter.GET("/captcha", admin.Lc.Captcha)
		/*******登录路由**********/
		adminRouter.GET("/login", admin.Lc.Login)
		adminRouter.POST("/login", admin.Lc.Login)
		adminRouter.GET("/login_out", admin.Lc.LoginOut)
		adminRouter.POST("/login_out", admin.Lc.LoginOut)

		adminHomeRouter := adminRouter.Group("/home")
		adminHomeRouter.Use(middleware.AdminUserAuth())
		{
			adminHomeRouter.GET("/", admin.Hc.Home)
			adminHomeRouter.GET("/welcome", admin.Hc.Welcome)
			adminHomeRouter.GET("/edit_password", admin.Hc.EditPassword)
			adminHomeRouter.POST("/save_password", admin.Hc.SavePassword)
			adminHomeRouter.POST("/save_skin", admin.Hc.SaveSkin)
		}

		adminSettingRouter := adminRouter.Group("/setting")
		adminSettingRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			adminGroup := adminSettingRouter.Group("/admingroup")
			{
				adminGroup.GET("/index", setting.Agc.Index)
				adminGroup.GET("/add", setting.Agc.AddIndex)
				adminGroup.POST("/save", setting.Agc.Save)
				adminGroup.GET("/edit", setting.Agc.Edit)
				adminGroup.GET("/del", setting.Agc.Del)
			}

			adminUser := adminSettingRouter.Group("/adminuser")
			{
				adminUser.GET("/index", setting.Auc.Index)
				adminUser.GET("/add", setting.Auc.AddIndex)
				adminUser.POST("/save", setting.Auc.Save)
				adminUser.GET("/edit", setting.Auc.Edit)
				adminUser.GET("/del", setting.Auc.Del)
			}

			adminSystem := adminSettingRouter.Group("/system")
			{
				adminSystem.GET("/index", setting.Asc.Index)
				adminSystem.GET("/getdir", setting.Asc.GetDir)
				adminSystem.GET("/views", setting.Asc.View)

				adminSystem.GET("/index_redis", setting.Asc.IndexRedis)
				adminSystem.GET("/getdir_redis", setting.Asc.GetDirRedis)
				adminSystem.GET("/view_redis", setting.Asc.ViewRedis)
			}
		}

		adminMinAppRouter := adminRouter.Group("/minApp")
		adminMinAppRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			minApp := adminMinAppRouter.Group("/minApp")
			{
				minApp.GET("/index", setting.Asc.Index)
			}
		}

	}
}
