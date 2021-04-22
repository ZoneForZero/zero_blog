package service

import (
	MODEL "zero_blog/model"
	SERIALIZER "zero_blog/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *SERIALIZER.Response {
	if service.PasswordConfirm != service.Password {
		return &SERIALIZER.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	MODEL.DB.Model(&MODEL.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &SERIALIZER.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}

	count = 0
	MODEL.DB.Model(&MODEL.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &SERIALIZER.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() SERIALIZER.Response {
	user := MODEL.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   MODEL.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return SERIALIZER.Err(
			SERIALIZER.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := MODEL.DB.Create(&user).Error; err != nil {
		return SERIALIZER.ParamErr("注册失败", err)
	}

	return SERIALIZER.BuildUserResponse(user)
}