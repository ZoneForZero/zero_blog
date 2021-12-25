package controller

import (
	"strconv"
	MODEL "zero_blog/model"
	SERIALIZER "zero_blog/serializer"
	// "github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

// 获取文章详情
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
	ctx.JSON(200, SERIALIZER.BuildArticleResponse(info, content))
}

// 获取文章列表（不用包含详情）
func GetArticles(ctx *gin.Context) {
	articles, err := MODEL.GetArticles()
	if err != nil {
		ctx.JSON(500, SERIALIZER.DBErr("文章列表获取异常!", err))
		return
	}
	ctx.JSON(200, SERIALIZER.Response{
		Msg: "获取列表成功",
		Data: SERIALIZER.BuildArticles(articles),
	})
}

// 添加文章列表
func AddArticle(ctx *gin.Context) {
	var param = new(struct {
		Title string `json:"title" binding:"required,max=30"`
		Content string `json:"content" binding:"required"`
	})
	if err := ctx.ShouldBindJSON(param); err != nil {
	 	ctx.JSON(433, SERIALIZER.DBErr("参数异常!", err))
	 	return
	}
	if err := MODEL.AddArticle(param.Title, param.Content, 1);err != nil {
		ctx.JSON(500, SERIALIZER.DBErr("添加失败!", err))
		return 
	}
	ctx.JSON(200, SERIALIZER.Response{
		Msg: "创建成功!",
		Data: param,
	})
}