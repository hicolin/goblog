package bootstrap

import (
	"embed"
	"goblog/pkg/view"
)

func SetupTemplate(tmplFS embed.FS) {
	view.TplFS = tmplFS
}
