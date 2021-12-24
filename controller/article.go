package controller

import (
	"strconv"
	MODEL "zero_blog/model"
	SERIALIZER "zero_blog/serializer"

	"github.com/gin-gonic/gin"
)

// UserMe 用户详情
func GetArticle(ctx *gin.Context) {
	articleIdString := ctx.Param("id")
	articleId, errCon := strconv.Atoi(articleIdString)
	if errCon != nil {
		ctx.JSON(433, SERIALIZER.ParamErr("id参数异常!", errCon))
		return
	}
	info, content, err1, err2 := MODEL.GetArticle(articleId)
	if err1 != nil {
		ctx.JSON(500, SERIALIZER.DBErr("文章信息获取异常!", err1))
		return
	} else if err2 != nil {
		ctx.JSON(500, SERIALIZER.DBErr("详细内容获取异常!", err2))
		return
	}
	res := SERIALIZER.BuildArticleResponse(info, content)
	ctx.JSON(200, res)
	return

}

// /wechat/applet_login?code=xxx [get]  路由
// 微信小程序登录
// func AppletWeChatLogin(ctx *gin.Context) {
// 	code := ctx.Query("code")         //  获取code
// 	userName := ctx.Query("userName") //  获取code

// 	// 根据code获取 openID 和 session_key
// 	wxLoginResp, err := USER_SERVICE.WXLogin(code, userName)
// 	if err != nil {
// 		ctx.JSON(400, ErrorResponse(err))
// 		return
// 	}
// 	// 保存登录态
// 	session := sessions.Default(ctx)
// 	session.Set("openid", wxLoginResp.OpenId)
// 	session.Set("sessionKey", wxLoginResp.SessionKey)
// 	// 这里用openid和sessionkey的串接 进行MD5之后作为该用户的自定义登录态
// 	mySession := GetMD5Encode(wxLoginResp.OpenId + wxLoginResp.SessionKey)
// 	// 接下来可以将openid 和 sessionkey, mySession 存储到数据库中,
// 	// 但这里要保证mySession 唯一, 以便于用mySession去索引openid 和sessionkey
// 	// c.String(100, mySession)
// 	ctx.JSON(200, mySession)
// }

// 将一个字符串进行MD5加密后返回加密后的字符串
// func GetMD5Encode(data string) string {
// 	h := md5.New()
// 	h.Write([]byte(data))
// 	return hex.EncodeToString(h.Sum(nil))
// }
