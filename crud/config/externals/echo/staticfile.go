package echo

import (
	"github.com/labstack/echo"
)

// setUpStaticFile 静的ファイルの設定を行う
// staticRootには静的ファイルノルートパスを指定する
func setUpStaticFile(staticRoot string, e *echo.Echo) {

	e.Static("/lib", staticRoot+"lib")
	e.Static("/index", staticRoot+"index.html")
}
