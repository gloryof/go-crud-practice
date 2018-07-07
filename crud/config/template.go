package config

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// Template テンプレートエンジンの設定
type Template struct {
	templates *template.Template
}

// CreateTemplate テンプレートエンジンの値を作成する
func CreateTemplate() Template {

	return Template{
		templates: template.Must(template.ParseGlob("crud/views/*.html")),
	}
}

// Render レンダリング処理（echo.Rendererの実装）
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
