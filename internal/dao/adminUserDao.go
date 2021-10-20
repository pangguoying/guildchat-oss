package dao

import (

	"gorm.io/gorm"
	"guildchat-oss/internal/models"
)

type adminUserDao struct {
	DB *gorm.DB
	Tx *gorm.DB
}

var AuDao = adminUserDao{DB: models.Db, Tx: models.Db.Begin()}
