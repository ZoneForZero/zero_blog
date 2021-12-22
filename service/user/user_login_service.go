package service

import (
	MODEL "zero_blog/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user MODEL.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// // Login 用户登录函数
// func (service *UserLoginService) Login(c *gin.Context) SERIALIZER.Response {
// 	var user MODEL.User

// 	if err := MODEL.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
// 		return SERIALIZER.ParamErr("账号或密码错误", nil)
// 	}

// 	if user.CheckPassword(service.Password) == false {
// 		return SERIALIZER.ParamErr("账号或密码错误", nil)
// 	}

// 	// 设置session
// 	service.setSession(c, user)

// 	return SERIALIZER.BuildUserResponse(user)
// }
