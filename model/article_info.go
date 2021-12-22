package model

import (
	GORM "github.com/jinzhu/gorm"
)

// User 用户模型
type ArticleInfo struct {
	GORM.Model
	AuthorId int    `gorm:"not null"`
	Title    string `gorm:"not null;size:255"`
	// Account        string	`gorm:"not null;primary_key;size:16"`
	// Password       string	`gorm:"not null"`
}

func GetArticlesByAuthor(authorId int) ([]ArticleInfo, error) {
	var article []ArticleInfo
	result := DB.Find(&article, ArticleInfo{AuthorId: authorId})
	return article, result.Error
}
