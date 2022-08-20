package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint   `gorm:"primaryKey" json:"id"`
	Email     string `gorm:"unique" validate:"required,email" json:"email"`
	FirstName string `gorm:"not null" validate:"required" json:"firstName"`
	LastName  string `gorm:"not null" validate:"required" json:"lastName"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"createdAt"`
	Updated   int64  `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}
