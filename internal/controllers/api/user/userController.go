package user

import (
	"guildchat-oss/internal/controllers/api"
	"guildchat-oss/internal/models"

	"github.com/gin-gonic/gin"
)

type userController struct {
	api.BaseController
}

var Uc = userController{}

func (apicon *userController) UserExample(c *gin.Context) {

	var (
		err     error
		userReq models.UserReq
	)
	err = apicon.FormBind(c, &userReq)

	if err != nil {
		apicon.Error(c, err)
		return
	}

	apicon.Success(c, userReq)
}
