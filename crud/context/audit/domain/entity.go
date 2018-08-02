package domain

import (
	"fmt"
	"time"
)

// AuditLog 監査ログ
type AuditLog struct {
	// operator 操作ユーザ
	operator Operator
	// operateAt 操作日時
	operateAt time.Time
	// operation 操作
	operation Operation
}

// GetOperator 操作ユーザを取得する
func (l AuditLog) GetOperator() Operator {

	return l.operator
}

// GetOperateAt 操作日時を取得する
func (l AuditLog) GetOperateAt() time.Time {

	return l.operateAt
}

// GetOperation 操作を取得する
func (l AuditLog) GetOperation() Operation {

	return l.operation
}

// CreateUserAuditLog ユーザ操作の監視ログを作成する。
// userIDにはユーザのIDを設定する。
// operateAtには操作日時を設定する。
// operationには操作内容の文字列を設定する。
func CreateUserAuditLog(userID uint64, operateAt time.Time, operation string) AuditLog {

	return AuditLog{
		operator:  registeredOperator{id: userID},
		operateAt: operateAt,
		operation: Operation{label: operation},
	}
}

// CreateAnnonymousAuditLog 匿名ユーザ操作の監視ログを作成する。
// operateAtには操作日時を設定する。
// operationには操作内容の文字列を設定する。
func CreateAnnonymousAuditLog(operateAt time.Time, operation string) AuditLog {

	return AuditLog{
		operator:  anonymousOperator{},
		operateAt: operateAt,
		operation: Operation{label: operation},
	}
}

// Operator 操作ユーザ
type Operator interface {
	// GetIdentifier 操作ユーザの識別子を取得する
	GetIdentifier() string
}

// registeredOperator 登録済みの操作ユーザ
type registeredOperator struct {
	id uint64
}

// GetIdentifier Operatorインターフェイスの実装
// IDを文字列で返す
func (o registeredOperator) GetIdentifier() string {

	return fmt.Sprint(o.id)
}

// anonymousOperator 匿名ユーザ
type anonymousOperator struct{}

// GetIdentifier Operatorインターフェイスの実装
// "anonymous"を返す
func (o anonymousOperator) GetIdentifier() string {

	return "anonymous"
}

// Operation 操作
type Operation struct {
	label string
}

// GetLabel ラベルを返す
func (o Operation) GetLabel() string {

	return o.label
}
