package admin

import (
	"github.com/gin-gonic/gin"
	"guildchat-oss/internal/controllers/admin"
)

type minAppController struct {
	admin.BaseController
}

var Mac = minAppController{}

func (con *minAppController) Index(c *gin.Context) {

}
