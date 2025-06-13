package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID uint   `gorm:"uniqueIndex"` // FK เชื่อมกับ User, 1:1
	Bio    string
	Avatar string
}
