package config

import (
	"errors"
	"io"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// LogConfig ログ設定
type LogConfig struct {
	// Access アクセスログ設定
	Access AccessLogConfig `json:"access"`
}

// AccessLogConfig アクセスログ設定
type AccessLogConfig struct {
	// Format フォーマット
	Format string `json:"format"`

	// Ouptu 出力先設定
	Output LogOutput `json:"output"`
}

// LogOutput ログ出力先設定
type LogOutput struct {
	// Type "Console" or "File"
	Type string `json:"type"`
	// Path 出力先パス
	Path string `json:"path"`
}

// SetUpLog ログの設定を行う
func SetUpLog(e *echo.Echo, param CrudParameter) error {

	lc, err := loadLogConfig(param)
	if err != nil {

		return err
	}

	e.Use(middleware.LoggerWithConfig(lc))

	return nil
}

// loadLogConfig ログの設定をロードする
func loadLogConfig(param CrudParameter) (middleware.LoggerConfig, error) {

	u := unmarshaller{}
	u.readFile(param.BasePath + "log.json")

	lf := defaultLogConfig()
	u.unmarshal(&lf)

	if u.err != nil {

		return middleware.LoggerConfig{}, u.err
	}

	w, we := getWriter(lf.Access)
	if we != nil {

		return middleware.LoggerConfig{}, we
	}

	lc := middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format:  lf.Access.Format,
		Output:  w,
	}

	return lc, nil
}

func defaultLogConfig() LogConfig {
	return LogConfig{
		Access: AccessLogConfig{
			Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
				`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
				`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
				`"bytes_out":${bytes_out}}` + "\n",
			Output: LogOutput{
				Type: "Console",
			},
		},
	}
}

func getWriter(config AccessLogConfig) (io.Writer, error) {

	switch config.Output.Type {
	case "Console":
		return os.Stdout, nil
	case "File":

		if f, err := os.Stat(config.Output.Path); os.IsNotExist(err) || !f.IsDir() {

			return os.Stderr, errors.New("ログのディレクトリが存在しません。[Path:" + config.Output.Path + "]")
		}

		logPath := config.Output.Path + "/access.log"
		file, fe := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if fe != nil {

			return os.Stderr, errors.New("ログファイルのオープンに失敗しました。[Path:" + logPath + "]")
		}

		return file, nil
	default:
		return os.Stderr, errors.New("Typeに不正な値が入力されました。")

	}
}
