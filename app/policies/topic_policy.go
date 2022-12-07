package policies

import (
	"goblog/app/http/models/article"
	"goblog/pkg/auth"
)

func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}
