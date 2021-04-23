package middleware

import (
	MODEL "zero_blog/model"
	SERIALIZER "zero_blog/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := MODEL.GetUserById(uid)
			if err == nil {
				ctx.Set("user", &user)
			}
		}
		ctx.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if user, _ := ctx.Get("user"); user != nil {
			if _, ok := user.(*MODEL.User); ok {
				ctx.Next()
				return
			}
		}

		ctx.JSON(200, SERIALIZER.CheckLogin())
		ctx.Abort()
	}
}
