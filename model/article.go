package model

import (
	GORM "github.com/jinzhu/gorm"
)

type ArticleInfo struct {
	GORM.Model
	AuthorId int    `gorm:"not null"`
	Title    string `gorm:"not null;size:255"`
	// Account        string	`gorm:"not null;primary_key;size:16"`
	// Password       string	`gorm:"not null"`
}
type ArticleContent struct {
	GORM.Model
	AtricleId int    `gorm:"not null"`
	Content   string `gorm:"not null;type:text"`
	// Account        string	`gorm:"not null;primary_key;size:16"`
	// Password       string	`gorm:"not null"`
}

// 获取列表，不包含文章详情
func GetArticlesByAuthor(authorId int) ([]ArticleInfo, error) {
	var article []ArticleInfo = make([]ArticleInfo, 0)
	result := DB.Find(&article, ArticleInfo{AuthorId: authorId})
	return article, result.Error
}

// 获取文章详情
func GetArticle(id int) (ArticleInfo, ArticleContent,error, error) {
	var articleInfo ArticleInfo
	var content ArticleContent
	infoResult := DB.First(&articleInfo, id)
	contentResult := DB.First(&ArticleContent, ArticleInfo{AtricleId: id})
	return infoResult, contentResult, infoResult.Error, contentResult.Error
}

func AddArticle(userId int, title string, content string) {
}