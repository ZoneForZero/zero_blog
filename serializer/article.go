package serializer

import MODEL "zero_blog/model"

type Article struct {
	ID        uint   `json:"id"`
	Title    string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"craeteTime"`
	UpdatedAt int64  `json:"updateTime"`
	DeletedAt int64  `json:"deleteTime"`
}

func BuildArticle(info MODEL.ArticleInfo, content MODEL.ArticleContent) Article {
	return Article{
		ID:       info.ID,
		Title: info.Title,
		Content: content.Content,
		CreatedAt: info.CreatedAt.Unix(),
	}
}

func BuildArticles(info []MODEL.ArticleInfo) []Article {
	//报错，会初始化为空，len不为0,但是append会在0后加 var result []Article = make([]Article, len(info))
	var result []Article
	for _, value := range info{
		result = append(result, Article {
			ID: value.ID,
			Title: value.Title,
			Content: "",
			CreatedAt: value.CreatedAt.Unix(),
		})
	}
	return result
}

func BuildArticleResponse(info MODEL.ArticleInfo, content MODEL.ArticleContent) Response {
	return Response{
		Data: BuildArticle(info,content),
	}
}
