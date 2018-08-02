package echo

import (
	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/config/registry"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CreateEcho Echoオブジェクトを作成する。
// logConfはログ設定を渡す。
// resultには依存性の登録結果を渡す。
func CreateEcho(logConf config.LogConfig, result *registry.Result) (*echo.Echo, error) {

	e := echo.New()
	e.Use(middleware.RequestID())

	if err := setUpLog(e, logConf, result); err != nil {

		return nil, err
	}

	setUpTemplate(e)
	setUpStaticFile(e)

	route(e, result)

	return e, nil
}

// Start サーバを起動する
func Start(e *echo.Echo) {

	e.Logger.Fatal(e.Start(":8000"))
}

// route ルーティングの設定を行う
func route(e *echo.Echo, result *registry.Result) {

	routeUser(e, result.User.Handler)
}