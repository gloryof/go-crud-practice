package domain

import (
	"github.com/gloryof/go-crud-practice/crud/context/base"
)

// newUserConvertedResult Userの変換結果
// 変換に成功した場合は下記の状態になる。
//   - name/birthDayに適切な値が設定されている
//   - errorsは空のエラーが設定されている
//
// 変換に失敗した場合は下記の状態になる。
//   - name/birthDayに初期構造体の値が設定される
//   - errorsはエラー含まれた状態が設定されている
type newUserConvertedResult struct {
	// name 名前
	name Name
	// birthDay 誕生日
	birthDay BirthDay
	// errors エラー結果
	errors base.ValidationResults
}

// newUserConvertedResultにIDが付与されたもの
// 変換に成功した場合は下記の状態になる。
//   - name/birthDay/idに適切な値が設定されている
//   - errorsは空のエラーが設定されている
//
// 変換に失敗した場合は下記の状態になる。
//   - name/birthDay/idに初期構造体の値が設定される
//   - errorsはエラー含まれた状態が設定されている
type existUserConvertedResult struct {
	// name 名前
	name Name
	// birthDay 誕生日
	birthDay BirthDay
	// id ID
	id ID
	// errors エラー結果
	errors base.ValidationResults
}

// RegisterSpec 新規登録時の仕様
type RegisterSpec struct {
	// Name 名前
	Name string
	// BirthDay 誕生日
	BirthDay string
}

// UpdateSpec 更新時の仕様
type UpdateSpec struct {
	// ID ID
	ID uint64
	// Name 名前
	Name string
	// BirthDay 誕生日
	BirthDay string
	// Repository リポジトリ
	Repository Repository
}

// Validate ユーザの入力値を検証する。
// 下記の入力検証が行われる
//   - 名前（domain.NewName）
//   - 誕生日（domain.BirthDay）
func (s RegisterSpec) Validate() base.ValidationResults {

	r := convertNewUserValues(s.Name, s.BirthDay)

	return r.errors
}

// Convert 新規ユーザに変換する
// 変換に失敗した場合はValidationResultsのポインタを返す。
func (s RegisterSpec) Convert() (User, *base.ValidationResults) {

	r := convertNewUserValues(s.Name, s.BirthDay)

	if r.errors.HasError() {

		return User{}, &r.errors
	}

	return User{
		name:     r.name,
		birthDay: r.birthDay,
	}, nil
}

// Validate ユーザの更新に対する入力値を検証する。
// 下記の入力検証が行われる。
//   - NewUseSpecと同等のチェック
//   - validateExistUser
func (s UpdateSpec) Validate() base.ValidationResults {

	r := convertExistUserValue(s.ID, s.Name, s.BirthDay, s.Repository)

	return r.errors
}

// Convert 更新内容をユーザに変換する。
// 変換に失敗した場合はValidationResultsのポインタを返す。
func (s UpdateSpec) Convert() (User, *base.ValidationResults) {

	r := convertExistUserValue(s.ID, s.Name, s.BirthDay, s.Repository)

	if r.errors.HasError() {

		return User{}, &r.errors
	}

	return User{
		id:         r.id,
		name:       r.name,
		birthDay:   r.birthDay,
		registered: true,
	}, nil
}

// validateExistUser 存在するユーザかを判定する
// 存在しない場合はエラーとして扱う
func validateExistUser(repository Repository, id ID) base.ValidationResults {

	r := base.NewValidationResults()

	_, er := repository.FindByID(id)
	if er != nil {

		r.AddDataNotExistsError("ユーザ")
	}

	return r
}

func convertNewUserValues(name string, birthDay string) newUserConvertedResult {

	r := newUserConvertedResult{
		name:     Name{},
		birthDay: BirthDay{},
		errors:   base.NewValidationResults(),
	}

	nv, ne := NewName(name)
	if ne != nil {

		r.errors.Add(*ne)
	} else {

		r.name = nv
	}

	bv, be := NewBirthDay(birthDay)
	if be != nil {

		r.errors.Add(*be)
	} else {

		r.birthDay = bv
	}

	return r
}

func convertExistUserValue(id uint64, name string, birthDay string, repository Repository) existUserConvertedResult {

	b := convertNewUserValues(name, birthDay)
	uid := ID{value: id}
	ee := validateExistUser(repository, uid)

	if ee.HasError() {

		b.errors.Add(ee)
	}

	return existUserConvertedResult{
		name:     b.name,
		birthDay: b.birthDay,
		errors:   b.errors,
		id:       ID{value: id},
	}
}
