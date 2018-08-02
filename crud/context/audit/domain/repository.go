package domain

// AuditRepository 監査リポジトリ
type AuditRepository interface {
	// Append ログを追記する
	Append(log AuditLog)
}
