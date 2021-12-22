package model

import (
	GORM "github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	GORM.Model
	OpenId   string `gorm:"not null;primary_key;size:64"`
	NickName string `gorm:"size:255"`
	// Account        string	`gorm:"not null;primary_key;size:16"`
	// Password       string	`gorm:"not null"`
}

// GetUserById 用ID获取用户
func GetUserById(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

func (userObj *User) UpsertUserByOpenId() error {
	var userResult User
	result := DB.First(&userResult, "open_id = ?", userObj.OpenId)
	// 找不到记录
	if result.Error != nil {
		// 创建用户
		if err := DB.Create(&userObj).Error; err != nil {
			return err
		}
	} else {
		// 更新用户
		if userResult.NickName != userObj.NickName {
			DB.Update(&userResult).Where("open_id = ?", userObj.OpenId).Update("nick_name", userObj.NickName)
		}
	}
	return nil
}

// SetPassword 设置密码
// func (user *User) SetPassword(password string) error {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
// 	if err != nil {
// 		return err
// 	}
// 	user.Password = string(bytes)
// 	return nil
// }
// // CheckPassword 校验密码
// func (user *User) CheckPassword(password string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 	return err == nil
// }
