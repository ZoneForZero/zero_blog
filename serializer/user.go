package serializer

import MODEL "zero_blog/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	NickName  string `json:"nick_name"`
	Account   string `json:"account"`
	OpenId    string `json:"open_id"`
	level     string `json:"level"`
	CreatedAt int64  `json:"created_at"`
}

// 转换函数    model.User  ->  User
func BuildUser(user MODEL.User) User {
	return User{
		ID:       user.ID,
		NickName: user.NickName,
		// Account:    user.Account,
		OpenId: user.OpenId,
		// Level:      user.Level,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// 序列化用户响应 model.User -> Response
func BuildUserResponse(user MODEL.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
