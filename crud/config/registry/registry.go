package registry

import (
	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/config/externals"
)

// Result 依存関係の登録結果
type Result struct {
	// User ユーザコンテキストの結果
	User *UserResult
	// Audit 監査コンテキストの結果
	Audit *AuditResult
}

// Register 依存関係の登録を行う
func Register(logConf config.LogConfig, context externals.Context) (Result, error) {

	ar, ae := registerAudit(logConf.Audit)
	if ae != nil {

		return Result{}, ae
	}

	ur := registerUser(context.DB.DBMap)
	return Result{
		Audit: &ar,
		User:  &ur,
	}, nil
}
