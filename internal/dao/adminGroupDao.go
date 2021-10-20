package dao

import (
	"gorm.io/gorm"
	"guildchat-oss/internal/models"
)

type adminGroupDao struct {
	DB *gorm.DB
}

var AgDao = adminGroupDao{DB: models.Db}
