package domain

import (
	"time"
	"unicode/utf8"

	"github.com/gloryof/go-crud-practice/crud/context/base"
)

// User ユーザを表すエンティティ
type User struct {
	// id ユーザのID
	id ID
	// name ユーザ名
	name Name
	// birthDay ユーザの誕生日
	birthDay BirthDay
	// registered 登録されたユーザ華道家のフラグ
	registered bool
}

// NewUser 新しいユーザを返す
func NewUser(id ID, name Name, birthDay BirthDay) User {

	return User{id: id, name: name, birthDay: birthDay}
}

// CreateExistUser 既存のユーザを返す
func CreateExistUser(id ID, name Name, birthDay BirthDay) User {

	return User{id: id, name: name, birthDay: birthDay, registered: true}
}

// IsExists 既存のユーザかどうかを判定する
// 既存ユーザの場合：true、既存ユーザではない場合：false
func (u User) IsExists() bool {
	return u.registered
}

// GetID ユーザIDを取得する
func (u User) GetID() ID { return u.id }

// GetName ユーザの名前を取得する
func (u User) GetName() Name { return u.name }

// GetBirthDay ユーザの誕生日を取得する
func (u User) GetBirthDay() BirthDay { return u.birthDay }

// ID ユーザを一意に特定するためのID
type ID struct {
	value uint64
}

// NewID 新しいIDの値を作成する
func NewID(value uint64) ID {

	return ID{value: value}
}

// GetValue ユーザIDの値を取得する
func (id ID) GetValue() uint64 {
	return id.value
}

// Name ユーザの名前を表す
type Name struct {
	value string
}

// NewName 新しいユーザの名前を作成する。
// 名前は下記の条件を満たす必要がある。
//  - 値が入力されている
//  - 40文字以内
// 条件を満たさない場合はValidationResultsを返す。
func NewName(value string) (Name, *base.ValidationResults) {

	const itemName = "名前"
	const maxLength = 40

	r := base.NewValidationResults()
	vl := utf8.RuneCountInString(value)
	if vl < 1 {

		r.AddRequiredError(itemName)
	}

	if maxLength < vl {

		r.AddMaxLengthError(itemName, maxLength)
	}

	if r.HasError() {

		return Name{}, &r
	}

	return Name{value: value}, nil
}

// GetValue ユーザ名の値を取得する
func (n Name) GetValue() string {
	return n.value
}

// BirthDay ユーザの誕生日
type BirthDay struct {
	value time.Time
}

// NewBirthDay 新しい誕生日の値を作成する
// 誕生日は下記の条件を満たす必要がある。
//  - 値が入力されている
//  - base.ConvertDateTimeの標準日付形式
// 条件を満たさない場合はValidationResultsを返す。
func NewBirthDay(value string) (BirthDay, *base.ValidationResults) {

	const itemName = "誕生日"

	r := base.NewValidationResults()
	if len(value) < 1 {

		r.AddRequiredError(itemName)
		return BirthDay{}, &r
	}

	cv, er := base.ConvertDateTime(value)
	if er != nil {

		r.AddDateFormatError(itemName)
		return BirthDay{}, &r
	}

	return BirthDay{value: cv}, nil
}

// ToBirthDay 誕生日に変換する
func ToBirthDay(value time.Time) BirthDay {

	return BirthDay{value: value}
}

// GetValue 誕生日の値を取得する
func (b BirthDay) GetValue() time.Time {
	return b.value
}
