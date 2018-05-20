package user

import (
	"time"

	"github.com/gloryof/go-crud-practice/context/base"
)

// User ユーザを表すエンティティ
type User struct {
	id       ID
	name     Name
	birthDay BirthDay
}

// NewUser 新しいユーザを作成する
func NewUser(id ID, name Name, birthDay BirthDay) User {
	return User{id, name, birthDay}
}

// GetID ユーザIDを取得する
func (u *User) GetID() ID { return u.id }

// GetName ユーザの名前を取得する
func (u *User) GetName() Name { return u.name }

// GetBirthDay ユーザの誕生日を取得する
func (u *User) GetBirthDay() BirthDay { return u.birthDay }

// ID ユーザを一意に特定するためのID
type ID struct {
	value uint64
}

// GetValue ユーザIDの値を取得する
func (id *ID) GetValue() uint64 {
	return id.value
}

// Name ユーザの名前を表す
type Name struct {
	value string
}

// GetValue ユーザ名の値を取得する
func (n *Name) GetValue() string {
	return n.value
}

// BirthDay ユーザの誕生日
type BirthDay struct {
	value time.Time
}

// GetValue 誕生日の値を取得する
func (b *BirthDay) GetValue() time.Time {
	return b.value
}

// Repository ユーザリポジトリ
type Repository interface {
	// FindById IDをキーにユーザを探す
	FindById(id ID) User

	// Save ユーザの保存を行う
	Save(user User)

	// Delete ユーザの削除を行う
	Delete(id ID)
}

// Convert ユーザに変換する
func Convert(name string, birthDay string) (User, *base.ValidationResults) {

	r := base.NewValidationResults()

	n, nv := convertName(name)

	if nv != nil {

		r = base.Merge(r, *nv)
	}

	b, bv := convertBirthDay(birthDay)
	if bv != nil {

		r = base.Merge(r, *bv)
	}

	if r.HasError() {

		return User{}, &r
	}

	return User{name: n, birthDay: b}, nil
}

// convertName 名前に変換する
// 変換に成功した場合：Name、変換に失敗した場合:base.ValidationResults
func convertName(name string) (Name, *base.ValidationResults) {

	const itemName = "名前"
	r := base.NewValidationResults()

	if len(name) < 1 {

		r.AddRequiredError(itemName)

		return Name{}, &r
	}

	return Name{name}, nil
}

// convertBirthday 誕生日に変換する
// 変換に成功した場合：BirthDay、変換に失敗した場合:base.ValidationResults
func convertBirthDay(birthDay string) (BirthDay, *base.ValidationResults) {

	const itemName = "誕生日"
	r := base.NewValidationResults()

	if len(birthDay) < 1 {

		r.AddRequiredError(itemName)
		return BirthDay{}, &r
	}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	b, br := time.ParseInLocation("2006-01-02", birthDay, loc)

	if br != nil {

		r.AddDateFormatError(itemName)
		return BirthDay{}, &r
	}

	return BirthDay{value: b}, nil
}
