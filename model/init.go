package model

import (
	"time"
	UTIL "zero_blog/util"

	GORM "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *GORM.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	db, err := GORM.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		UTIL.Log().Panic("连接数据库不成功", err)
	} else {
		UTIL.Log().Println("连接Mysql成功!")
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	// 全局禁用表名复数
	db.SingularTable(true)
	DB = db
	migration()
}

//执行数据迁移
func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&ArticleInfo{})
	DB.AutoMigrate(&ArticleContent{})
}
