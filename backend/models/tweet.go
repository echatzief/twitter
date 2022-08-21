package models

import "gorm.io/gorm"

type Tweet struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	UserId    uint           `json:"userId" gorm:"column:user_id"`
	Tweet     string         `json:"tweet" gorm:"unique"`
	CreatedAt int64          `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt int64          `json:"updatedAt" gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
