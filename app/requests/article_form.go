package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/http/models/article"
)

func ValidateArticleForm(data article.Article) map[string][]string {
	rules := govalidator.MapData{
		"title": []string{"required", "min:3", "max:40"},
		"body":  []string{"required", "min:10"},
	}

	messages := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
			"min:标题长度需大于 3",
			"max:标题长度需大于 10",
		},
		"body": []string{
			"required:文章内容为必填项",
			"min:长度需大于 10",
		},
	}

	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
