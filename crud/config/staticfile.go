package config

import (
	"github.com/labstack/echo"
)

// SetUpStaticFile 静的ファイルの設定を行う
func SetUpStaticFile(e *echo.Echo) {

	e.Static("/lib", "./static/lib")
	e.Static("/index", "./static/index.html")
}
