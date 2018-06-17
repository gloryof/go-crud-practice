package domain

import (
	"time"
)

// User ユーザを表すエンティティ
type User struct {
	id       ID
	name     Name
	birthDay BirthDay
}

// IsExists 既存のユーザかどうかを判定する
// 既存ユーザの場合：true、既存ユーザではない場合：false
func (u User) IsExists() bool {
	return u.id.numbered
}

// GetID ユーザIDを取得する
func (u User) GetID() ID { return u.id }

// GetName ユーザの名前を取得する
func (u User) GetName() Name { return u.name }

// GetBirthDay ユーザの誕生日を取得する
func (u User) GetBirthDay() BirthDay { return u.birthDay }

// ID ユーザを一意に特定するためのID
type ID struct {
	numbered bool
	value    uint64
}

// GetValue ユーザIDの値を取得する
func (id ID) GetValue() uint64 {
	return id.value
}

// Name ユーザの名前を表す
type Name struct {
	value string
}

// GetValue ユーザ名の値を取得する
func (n Name) GetValue() string {
	return n.value
}

// BirthDay ユーザの誕生日
type BirthDay struct {
	value time.Time
}

// GetValue 誕生日の値を取得する
func (b BirthDay) GetValue() time.Time {
	return b.value
}
