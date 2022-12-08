package config

import "goblog/pkg/config"

func init() {
	config.Add("pagination", config.StrMap{
		"perpage":   10,
		"url_query": "page",
	})
}
