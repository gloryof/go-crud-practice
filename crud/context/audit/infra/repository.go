package infra

import (
	"github.com/gloryof/go-crud-practice/crud/context/audit/domain"
	"github.com/gloryof/go-crud-practice/crud/externals/zap/audit"
)

// AuditRepositoryLoggerImpl 監査リポジトリのロガー実装
type AuditRepositoryLoggerImpl struct {
	// logger ロガー
	logger *audit.Logger
}

// NewAuditRepositoryLoggerImpl AuditRepositoryLoggerImplを生成する
func NewAuditRepositoryLoggerImpl(logger *audit.Logger) AuditRepositoryLoggerImpl {

	return AuditRepositoryLoggerImpl{
		logger: logger,
	}
}

// Append 監視リポジトリの実装
func (a AuditRepositoryLoggerImpl) Append(log domain.AuditLog) {

	a.logger.RecordAudit(log)
}
