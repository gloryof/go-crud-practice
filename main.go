package main

import (
	"errors"
	"flag"
	"os"
	"strings"

	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/config/externals"
	"github.com/gloryof/go-crud-practice/crud/config/externals/echo"
	"github.com/gloryof/go-crud-practice/crud/config/registry"
	"github.com/labstack/gommon/log"
)

var (
	paramC = flag.String("c", "./config/development/", "Config base directory")
	paramS = flag.String("s", "./static/", "Config base directory")
)

// main 起動処理
// -c 設定ファイルのディレクトリパス。デフォルトは実行ファイルがあるディレクトリ内のconfigディレクトリ。
// -s 静的ファイルのルートディレクトリパス。デフォルトはカレントディレクトリ。
func main() {

	flag.Parse()

	p, pe := loadParameter()
	if pe != nil {

		handleError(pe)
	}

	lc, le := config.LoadLogConfig(p)
	if le != nil {

		handleError(le)
	}

	// TODO DBの設定は入れたけどトランザクション制御のテストをしていないので後ほど確認する
	dc, de := config.LoadDBConfig(p)
	if de != nil {

		handleError(de)
	}

	ctx, err := externals.CreateContext(dc)
	if err != nil {

		handleError(err)
	}

	defer ctx.Close()

	rr, re := registry.Register(lc, ctx)
	if re != nil {

		handleError(re)
	}

	e, ee := echo.CreateEcho(p, lc, &rr)
	if ee != nil {

		handleError(ee)
	}
	echo.Start(e)
}

func loadParameter() (config.CrudParameter, error) {

	p := loadBasePath()

	if f, err := os.Stat(p); os.IsNotExist(err) || !f.IsDir() {

		return config.CrudParameter{}, errors.New("設定ファイルのディレクトリが存在しません[" + p + "]")
	}

	r := loadStaticRoot()

	if f, err := os.Stat(r); os.IsNotExist(err) || !f.IsDir() {

		return config.CrudParameter{}, errors.New("静的ファイルのルートディレクトリが存在しません[" + r + "]")
	}

	return config.CrudParameter{
		BasePath:            p,
		StaticRootDirectory: r,
	}, nil
}

func loadBasePath() string {

	v := *paramC
	if strings.HasSuffix(v, "/") {

		return v
	}

	return v + "/"
}

func loadStaticRoot() string {

	v := *paramS
	if strings.HasSuffix(v, "/") {

		return v
	}

	return v + "/"
}

func handleError(err error) {

	log.Errorf("Error!: %+v\n", err)
	os.Exit(1)
}
