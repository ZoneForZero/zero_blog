package model

import (
	GORM "github.com/jinzhu/gorm"
)

// User 用户模型
type ArticleContent struct {
	GORM.Model
	AtricleId int    `gorm:"not null"`
	Content   string `gorm:"not null;type:text"`
	// Account        string	`gorm:"not null;primary_key;size:16"`
	// Password       string	`gorm:"not null"`
}
