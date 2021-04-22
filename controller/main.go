package controller

import (
	"encoding/json"
	"fmt"
	CONF "zero_blog/conf"
	MODEL "zero_blog/model"
	SERIALIZER "zero_blog/serializer"
	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
)

// Ping 状态检查页面
func Ping(ctx *gin.Context) {
	ctx.JSON(200, SERIALIZER.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(ctx *gin.Context) *MODEL.User {
	if user, _ := ctx.Get("user"); user != nil {
		if u, ok := user.(*MODEL.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) SERIALIZER.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := CONF.T(fmt.Sprintf("Field.%s", e.Field))
			tag := CONF.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return SERIALIZER.ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return SERIALIZER.ParamErr("JSON类型不匹配", err)
	}

	return SERIALIZER.ParamErr("参数错误", err)
}
