package controller

import (
	"crypto/md5"
	"encoding/hex"
	SERIALIZER "zero_blog/serializer"
	USER_SERVICE "zero_blog/service/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
// func UserRegister(ctx *gin.Context) {
// 	var service USER_SERVICE.UserRegisterService
// 	if err := ctx.ShouldBind(&service); err == nil {
// 		res := service.Register()
// 		ctx.JSON(200, res)
// 	} else {
// 		ctx.JSON(200, ErrorResponse(err))
// 	}
// }

// UserLogin 用户登录接口
// func UserLogin(ctx *gin.Context) {
// 	var service USER_SERVICE.UserLoginService
// 	if err := ctx.ShouldBind(&service); err == nil {
// 		res := service.Login(ctx)
// 		ctx.JSON(200, res)
// 	} else {
// 		ctx.JSON(200, ErrorResponse(err))
// 	}
// }

// UserMe 用户详情
func UserMe(ctx *gin.Context) {
	user := CurrentUser(ctx)
	// user Model对象转换成用户数据
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

// /wechat/applet_login?code=xxx [get]  路由
// 微信小程序登录
func AppletWeChatLogin(ctx *gin.Context) {
	code := ctx.Query("code")         //  获取code
	userName := ctx.Query("userName") //  获取code

	// 根据code获取 openID 和 session_key
	wxLoginResp, err := USER_SERVICE.WXLogin(code, userName)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		return
	}
	// 保存登录态
	session := sessions.Default(ctx)
	session.Set("openid", wxLoginResp.OpenId)
	session.Set("sessionKey", wxLoginResp.SessionKey)
	// 这里用openid和sessionkey的串接 进行MD5之后作为该用户的自定义登录态
	mySession := GetMD5Encode(wxLoginResp.OpenId + wxLoginResp.SessionKey)
	// 接下来可以将openid 和 sessionkey, mySession 存储到数据库中,
	// 但这里要保证mySession 唯一, 以便于用mySession去索引openid 和sessionkey
	// c.String(100, mySession)
	ctx.JSON(200, mySession)
}

// 将一个字符串进行MD5加密后返回加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
