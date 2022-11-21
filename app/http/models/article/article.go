package article

import (
	"goblog/app/http/models"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel

	Title string
	Body  string
}

func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}
