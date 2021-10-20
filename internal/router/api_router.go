package router

import (
	"github.com/gin-gonic/gin"
	"guildchat-oss/internal/controllers/api/user"
)

func ApiRouter(apiRouter *gin.RouterGroup) {
	apiRouter.Use()
	{
		apiUserRouter := apiRouter.Group("user")
		{
			apiUserRouter.POST("/example", user.Uc.UserExample)
		}
	}
}
