package models

import "gorm.io/gorm"

//user model

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Role     string `gorm:"not null;default:user" json:"role"`
}
