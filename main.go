package main

import (
	"errors"
	"flag"
	"os"
	"strings"

	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/config/router"
	"github.com/gloryof/go-crud-practice/crud/externals"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var (
	paramC = flag.String("c", "./config/", "Config base directory")
)

// main 起動処理
// -c 設定ファイルのディレクトリパス。デフォルトは実行ファイルがあるディレクトリ内のconfigディレクトリ。
func main() {

	flag.Parse()

	c, ce := loadParameter()
	if ce != nil {

		handleError(ce)
	}

	e := createEcho()

	ctx, err := createExternalsContext(c)
	if err != nil {

		handleError(err)
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

		return externals.Context{}, ce
	}

	ctx, cte := externals.CreateContext(c)

	if cte != nil {

		ctx.Close()
		return externals.Context{}, cte
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

func loadParameter() (config.CrudParameter, error) {

	p := loadBasePath()

	if f, err := os.Stat(p); os.IsNotExist(err) || !f.IsDir() {

		return config.CrudParameter{}, errors.New("設定ファイルのディレクトリが存在しません")
	}

	return config.CrudParameter{
		BasePath: p,
	}, nil
}

func loadBasePath() string {

	v := *paramC
	if strings.HasSuffix(v, "/") {

		return v
	}

	return v + "/"
}

func handleError(err error) {

	log.Errorf("Error!: %s", err)
	os.Exit(1)
}
