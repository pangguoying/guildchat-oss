package admin

import (
	"github.com/gin-gonic/gin"
	"guildchat-oss/internal/controllers/admin"
)

type minAppController struct {
	admin.BaseController
}

var Mc = minAppController{}

func (con *minAppController) Index(c *gin.Context) {

}
