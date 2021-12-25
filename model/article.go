package model

import (
	GORM "github.com/jinzhu/gorm"
	// "fmt"
)

type ArticleInfo struct {
	GORM.Model
	AuthorId int    `gorm:"not null"`
	Title    string `gorm:"not null;size:30"`
	// Account        string	`gorm:"not null;primary_key;size:16"`
	// Password       string	`gorm:"not null"`
}
type ArticleContent struct {
	GORM.Model
	ArticleId int    `gorm:"not null"`
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
func GetArticle(id int) (ArticleInfo, ArticleContent, error, error) {
	var articleInfo ArticleInfo
	var content ArticleContent
	infoResult := DB.First(&articleInfo, id)
	contentResult := DB.First(&content, ArticleContent{ArticleId: id})
	return articleInfo, content, infoResult.Error, contentResult.Error
}

// 获取文章列表
func GetArticles() ([]ArticleInfo, error) {
	var articles []ArticleInfo
	articlesDbResult := DB.Find(&articles)
	return articles, articlesDbResult.Error
}


func AddArticle(title string, content string, userId int) error {
	var infoDbObj = ArticleInfo {
		AuthorId: userId,
		Title: title,
	}
	var contentDbObj = ArticleContent {
		ArticleId: 0,
		Content:content,
	}
	// 找不到记录
	if err := DB.Create(&infoDbObj).Error; err != nil {
		return err
	}
	// 取創建的id進行關聯
	var infoObj ArticleInfo
	createResult := DB.First(&infoObj, infoDbObj)
	if createResult.Error != nil {
		return createResult.Error
	}
	contentDbObj.ArticleId = int(infoObj.ID)
	if err := DB.Create(&contentDbObj).Error; err != nil {
		// 回滾删除info数据
		return err
	}
	return nil
}


// func AddArticle(userId int, title string, content string) {
// }
