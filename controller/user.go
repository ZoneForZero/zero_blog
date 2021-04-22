package controller

import (
	SERIALIZER "zero_blog/serializer"
	USER_SERVICE "zero_blog/service/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(ctx *gin.Context) {
	var service USER_SERVICE.UserRegisterService
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.Register()
		ctx.JSON(200, res)
	} else {
		ctx.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(ctx *gin.Context) {
	var service USER_SERVICE.UserLoginService
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.Login(ctx)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(ctx *gin.Context) {
	user := CurrentUser(ctx)
	res := SERIALIZER.BuildUserResponse(*user)
	ctx.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(ctx *gin.Context) {
	s := sessions.Default(ctx)
	s.Clear()
	s.Save()
	ctx.JSON(200, SERIALIZER.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
