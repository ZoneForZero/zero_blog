package service

import (
	MODEL "zero_blog/model"
	SERIALIZER "zero_blog/serializer"
)

// 用户传参格式校验
type UserRegisterService struct {
	OpenId          string `form:"open_id" json:"open_id" binding:"min=5,max=30"`
	Account         string `form:"account" json:"account" binding:"min=5,max=30"`
	NickName        string `form:"nick_name" json:"nick_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// 用户传参逻辑校验
func (service *UserRegisterService) valid() *SERIALIZER.Response {
	if service.PasswordConfirm != service.Password {
		return &SERIALIZER.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	// 账号和名称唯一
	checkFieldNames:= []string{"account", "nick_name"}
	checkValue := []string{service.Account,service.NickName}
	checkLength := 2
	for i := 0; i < checkLength; i++ {
		count := 0
		preString := checkFieldNames[i] + " = ?"
		MODEL.DB.Model(&MODEL.User{}).Where(preString, checkValue[i]).Count(&count)
		if count > 0 {
			return &SERIALIZER.Response{
				Code: 40001,
				Msg:  checkFieldNames[i] + "被占用",
			}
		}
	}
	// count = 0
	// MODEL.DB.Model(&MODEL.User{}).Where("user_name = ?", service.UserName).Count(&count)
	// if count > 0 {
	// 	return &SERIALIZER.Response{
	// 		Code: 40001,
	// 		Msg:  "用户名已经注册",
	// 	}
	// }

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() SERIALIZER.Response {
	user := MODEL.User{
		NickName: service.NickName,
		Account: service.Account,
		OpenId: service.OpenId,
		Level: 1,
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