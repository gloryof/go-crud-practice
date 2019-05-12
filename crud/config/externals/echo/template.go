package echo

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template テンプレートエンジンの設定
type Template struct {
	templates *template.Template
}

// setUpTemplate テンプレートエンジンの設定を行う
// rootPathにはテンプレートエンジンのルートとなるパスを設定する
func setUpTemplate(rootPath string, e *echo.Echo) {

	e.Renderer = Template{
		templates: template.Must(template.ParseGlob(rootPath + "views/**/*.html")),
	}
}

// Render レンダリング処理（echo.Rendererの実装）
func (t Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
