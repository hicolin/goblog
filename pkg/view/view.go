package view

import (
	"goblog/app/http/models/category"
	"goblog/app/http/models/user"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type D map[string]interface{}

func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {
	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User()
	data["flash"] = flash.All()
	data["Users"], _ = user.All()
	data["Categories"], _ = category.All()

	allFiles := getTemplateFiles(tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}

func getTemplateFiles(tplFiles ...string) []string {
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	return append(layoutFiles, tplFiles...)

}
