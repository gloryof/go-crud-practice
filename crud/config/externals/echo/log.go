package echo

import (
	"errors"
	"io"
	"os"

	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/config/registry"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// setUpLog ログの設定を行う
func setUpLog(e *echo.Echo, conf config.LogConfig, result *registry.Result) error {

	lc, err := loadMiddlewareLogConfig(conf)
	if err != nil {

		return err
	}

	e.Use(middleware.LoggerWithConfig(lc))
	e.Use(createAuditLogFunc(result))

	return nil
}

// loadMiddlewareLogConfig ミドルウェアのログの設定をロードする
func loadMiddlewareLogConfig(conf config.LogConfig) (middleware.LoggerConfig, error) {

	w, we := getWriter(conf.Access)
	if we != nil {

		return middleware.LoggerConfig{}, we
	}

	lc := middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format:  conf.Access.Format,
		Output:  w,
	}

	return lc, nil
}

func getWriter(conf config.AccessLogConfig) (io.Writer, error) {

	switch conf.Output.Type {
	case config.OutputConsole:
		return os.Stdout, nil
	case config.OutputFile:

		if f, err := os.Stat(conf.Output.Path); os.IsNotExist(err) || !f.IsDir() {

			return os.Stderr, errors.New("ログのディレクトリが存在しません。[Path:" + conf.Output.Path + "]")
		}

		logPath := conf.Output.Path + "/access.log"
		file, fe := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if fe != nil {

			return os.Stderr, errors.New("ログファイルのオープンに失敗しました。[Path:" + logPath + "]")
		}

		return file, nil
	default:
		return os.Stderr, errors.New("Typeに不正な値が入力されました。")

	}
}

func createAuditLogFunc(result *registry.Result) echo.MiddlewareFunc {

	a := result.Audit.Usecase

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			a.Audit.RecordAnnonymousOperation(c.Path())
			err := next(c)
			return err
		}
	}
}
