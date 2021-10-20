//+build !embed

/*
 * @Description:部署拷贝静态文件时忽略该文件
 * @Author: pgy
 * @Date: 2021-10-20 10:59:00
 */
package web

import (
	"guildchat-oss/pkg/comment"
	"net/http"
	"path/filepath"
	"strings"

	template2 "guildchat-oss/pkg/template"

	"github.com/gin-contrib/multitemplate"
)

var rootPath string
var StaticsFs http.FileSystem

func init() {
	rootPath, _ = comment.RootPath()

	StaticsFs = http.Dir(rootPath + string(filepath.Separator) + "web" + string(filepath.Separator) + "statics")
}

func LoadTemplates() multitemplate.Renderer {
	templatesDir := rootPath + "/web/views"
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layout/*.html")

	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/template/*/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		dirSlice := strings.Split(include, string(filepath.Separator))
		fileName := strings.Join(dirSlice[len(dirSlice)-2:], "/")
		r.AddFromFilesFuncs(fileName, template2.GlobalTemplateFun, files...)
	}
	return r
}
