package echo

import (
	"fmt"

	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/config/registry"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CreateEcho Echoオブジェクトを作成する。
// crudParamにはアプリのパラメータを設定する
// logConfはログ設定を渡す。
// resultには依存性の登録結果を渡す。
func CreateEcho(crudParam config.CrudParameter, logConf config.LogConfig, result *registry.Result) (*echo.Echo, error) {

	e := echo.New()
	e.Use(middleware.RequestID())

	if err := setUpLog(e, logConf, result); err != nil {

		return nil, err
	}

	setUpTemplate(crudParam.StaticRootDirectory, e)
	setUpStaticFile(crudParam.StaticRootDirectory, e)

	route(e, result)

	return e, nil
}

// Start サーバを起動する
func Start(e *echo.Echo, port int) {

	e.Logger.Fatal(e.Start(":" + fmt.Sprint(port)))
}

// route ルーティングの設定を行う
func route(e *echo.Echo, result *registry.Result) {

	routeUser(e, result.User.Handler)
}
