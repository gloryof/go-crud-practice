package usecase

import (
	"time"

	"github.com/gloryof/go-crud-practice/crud/context/audit/domain"
)

// Audit 監査ユースケース
type Audit struct {
	// clockGenerator 時刻生成関数
	clockGenerator func() time.Time
	// repository 監査リポジトリ
	repository domain.AuditRepository
}

// NewAudit Auditoを生成する。
// genrateには時刻を生成する関数を設定する。
// repositoryには監査リポジトリを設定する。
func NewAudit(clockGenerator func() time.Time, repository domain.AuditRepository) Audit {
	return Audit{
		clockGenerator: clockGenerator,
		repository:     repository,
	}
}

// RecordUserOperation ユーザによる操作を記録する。
// userIDは操作したユーザのIDを設定する。
// operationには操作した内容の文字列を設定する。
func (a Audit) RecordUserOperation(userID uint64, operation string) {

	l := domain.CreateUserAuditLog(userID, a.clockGenerator(), operation)

	a.repository.Append(l)
}

// RecordAnnonymousOperation 匿名ユーザによる操作を記録する。
// operationには操作した内容の文字列を設定する。
func (a Audit) RecordAnnonymousOperation(operation string) {

	l := domain.CreateAnnonymousAuditLog(a.clockGenerator(), operation)

	a.repository.Append(l)
}
