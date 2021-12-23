package serializer

import MODEL "zero_blog/model"

// User 用户序列化器
type Article struct {
	ID        uint   `json:"id"`
	Title    string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"craeteTime"`
}

// 转换函数    model.User  ->  User
func BuildArticle(info MODEL.ArticleInfo, content MODEL.ArticleContent) Article {
	return Article{
		ID:       info.ID,
		Title: info.Title,
		Content: content.Content,
		CreatedAt: info.CreatedAt.Unix(),
	}
}

// 序列化用户响应 model.User -> Response
func BuildArticleResponse(info MODEL.ArticleInfo, content MODEL.ArticleContent) Response {
	return Response{
		Data: BuildArticle(info,content),
	}
}
