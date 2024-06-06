package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body   string `gorm:"tye:varchar(255)" json:"body"`
	UserID uint
	PostID uint
}
