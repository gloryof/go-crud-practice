package echo

import (
	"github.com/labstack/echo"
)

// setUpStaticFile 静的ファイルの設定を行う
func setUpStaticFile(e *echo.Echo) {

	e.Static("/lib", "./static/lib")
	e.Static("/index", "./static/index.html")
}
