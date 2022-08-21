package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"unique" validate:"required"`
	Password  string         `json:"password" validate:"required"`
	Email     string         `json:"email" gorm:"unique" validate:"required,email"`
	FirstName string         `json:"firstName" gorm:"not null" validate:"required"`
	LastName  string         `json:"lastName" gorm:"not null" validate:"required"`
	CreatedAt int64          `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt int64          `json:"updatedAt" gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Hooks
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := HashPassword(u.Password)
		if err != nil {
			return nil
		}
		u.Password = hash
	}
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := HashPassword(u.Password)
		if err != nil {
			return nil
		}
		u.Password = hash
	}
	return
}
