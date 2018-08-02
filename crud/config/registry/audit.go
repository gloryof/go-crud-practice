package registry

import (
	"time"

	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/context/audit/infra"
	"github.com/gloryof/go-crud-practice/crud/context/audit/usecase"
	"github.com/gloryof/go-crud-practice/crud/externals/zap/audit"
)

// AuditResult Auditのregistryの結果
type AuditResult struct {
	// Infra インフラレイヤ
	Infra *AuditInfra
	// Usecase ユースケースレイヤ
	Usecase *AuditUsecase
}

// AuditInfra Auditのインフラ群
type AuditInfra struct {
	// Repository リポジトリ
	Repository *infra.AuditRepositoryLoggerImpl
}

// AuditUsecase Auditのユースケース群
type AuditUsecase struct {
	// Audit 監査ユースケース
	Audit *usecase.Audit
}

// registerAudit Auditの依存関係を登録する
func registerAudit(logConf config.AuditLogConfig) (AuditResult, error) {

	ir, ie := registerAuditInfra(logConf)
	if ie != nil {

		return AuditResult{}, ie
	}

	u := registerAuditUsecase(&ir)
	return AuditResult{
		Infra:   &ir,
		Usecase: &u,
	}, nil
}

// registerAuditInfra 依存性の登録を行う
func registerAuditInfra(logConf config.AuditLogConfig) (AuditInfra, error) {

	l, e := audit.NewLogger(logConf)

	if e != nil {

		return AuditInfra{}, e
	}

	i := infra.NewAuditRepositoryLoggerImpl(&l)
	return AuditInfra{
		Repository: &i,
	}, nil
}

// registerAuditUsecase 依存性の登録を行う
func registerAuditUsecase(infra *AuditInfra) AuditUsecase {
	a := usecase.NewAudit(time.Now, infra.Repository)
	return AuditUsecase{
		Audit: &a,
	}
}
