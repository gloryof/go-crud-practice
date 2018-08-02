package audit

import (
	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/context/audit/domain"
	"go.uber.org/zap"
)

// Logger 監査ロガー
type Logger struct {
	logger *zap.Logger
}

// NewLogger 監査ロガーを作成する。
// logConfは監査ログの設定。
func NewLogger(logConf config.AuditLogConfig) (Logger, error) {

	c := zap.NewProductionConfig()

	switch logConf.Output.Type {
	case config.OutputConsole:
		c.OutputPaths = []string{"stdout"}
	case config.OutputFile:
		c.OutputPaths = []string{logConf.Output.Path}
	default:
		c.OutputPaths = []string{"stdout"}
	}

	l, err := c.Build()

	if err != nil {

		return Logger{}, err
	}

	return Logger{
		logger: l,
	}, nil
}

// RecordAudit 監査ログを記録する
func (l Logger) RecordAudit(log domain.AuditLog) {

	lo := l.logger

	lo.Info("Record audit log.",
		zap.Time("operateAt", log.GetOperateAt()),
		zap.String("operator", log.GetOperator().GetIdentifier()),
		zap.String("operation", log.GetOperation().GetLabel()),
	)
}
