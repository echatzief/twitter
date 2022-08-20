package models

import (
	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	Id        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"column:user_id"`
	Tweet     string `gorm:"unique"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	Updated   int64  `gorm:"autoUpdateTime:milli"`
}
