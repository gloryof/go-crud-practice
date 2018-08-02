package config

// OutputConsole 出力タイプ：コンソール
const OutputConsole string = "Console"

// OutputFile 出力タイプ：ファイル
const OutputFile string = "File"

// LogConfig ログ設定
type LogConfig struct {
	// Access アクセスログ設定
	Access AccessLogConfig `json:"access"`
	// Audit 監査ログ設定
	Audit AuditLogConfig `json:"audit"`
}

// AccessLogConfig アクセスログ設定
type AccessLogConfig struct {
	// Format フォーマット
	Format string `json:"format"`

	// Ouptu 出力先設定
	Output LogOutput `json:"output"`
}

// AuditLogConfig 監査ログ設定
type AuditLogConfig struct {

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

// LoadLogConfig ログの設定をロードする
func LoadLogConfig(param CrudParameter) (LogConfig, error) {

	u := unmarshaller{}
	u.readFile(param.BasePath + "log.json")

	lf := defaultLogConfig()
	u.unmarshal(&lf)

	if u.err != nil {

		return LogConfig{}, u.err
	}

	return lf, nil
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
		Audit: AuditLogConfig{
			Output: LogOutput{
				Type: "Console",
			},
		},
	}
}
