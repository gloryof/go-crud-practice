package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/config/router"
	"github.com/gloryof/go-crud-practice/crud/externals"
	"github.com/labstack/echo"
)

var (
	paramC = flag.String("c", "./config/", "Config base directory")
)

// main 起動処理
// -c 設定ファイルのディレクトリパス。デフォルトは実行ファイルがあるディレクトリ内のconfigディレクトリ。
func main() {

	flag.Parse()

	c := loadParameter()
	e := createEcho()

	ctx, err := createExternalsContext(c)
	if err != nil {

		os.Exit(1)
	}

	defer ctx.Close()

	route(e, ctx)

	start(e)
}

func createEcho() *echo.Echo {

	e := echo.New()
	config.SetUpTemplate(e)
	config.SetUpStaticFile(e)

	return e
}

func createExternalsContext(p config.CrudParameter) (externals.Context, error) {

	c, ce := config.LoadDBConfig(p)

	if ce != nil {

		return externals.Context{}, fmt.Errorf("Error!: %s", ce)
	}

	ctx, cte := externals.CreateContext(c)

	if cte != nil {

		ctx.Close()
		return externals.Context{}, fmt.Errorf("Error!: %s", cte)
	}

	return ctx, nil
}

func route(e *echo.Echo, ctx externals.Context) {

	ug := e.Group("/user")
	router.User(ctx, ug)

}

func start(e *echo.Echo) {

	e.Start(":8000")
}

func loadParameter() config.CrudParameter {

	return config.CrudParameter{
		BasePath: loadBasePath(),
	}
}

func loadBasePath() string {

	v := *paramC
	if strings.HasSuffix(v, "/") {

		return v
	}

	return v + "/"
}
