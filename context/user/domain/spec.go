package domain

import (
	"time"

	"github.com/gloryof/go-crud-practice/context/base"
)

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
//   - validateName
//   - validateBirthDay
func (s RegisterSpec) Validate() base.ValidationResults {

	r := base.NewValidationResults()

	r.Add(validateName(s.Name))
	r.Add(validateBirthDay(s.BirthDay))

	return r
}

// Convert 新規ユーザに変換する
// 変換に失敗した場合はValidationResultsのポインタを返す。
func (s RegisterSpec) Convert() (User, *base.ValidationResults) {

	r := s.Validate()
	if r.HasError() {

		return User{}, &r
	}

	n := convertName(s.Name)
	b := convertBirthDay(s.BirthDay)

	return User{name: n, birthDay: b}, nil
}

// Validate ユーザの更新に対する入力値を検証する。
// 下記の入力検証が行われる。
//   - NewUseSpecと同等のチェック
//   - validateExistUser
func (s UpdateSpec) Validate() base.ValidationResults {

	uid := ID{value: s.ID}
	r := RegisterSpec{Name: s.Name, BirthDay: s.BirthDay}.Validate()

	r.Add(validateExistUser(s.Repository, uid))

	return r
}

// Convert 更新内容をユーザに変換する。
// 変換に失敗した場合はValidationResultsのポインタを返す。
func (s UpdateSpec) Convert() (User, *base.ValidationResults) {

	r := s.Validate()
	if r.HasError() {

		return User{}, &r
	}

	id := ID{numbered: true, value: s.ID}
	n := convertName(s.Name)
	b := convertBirthDay(s.BirthDay)

	return User{id: id, name: n, birthDay: b}, nil
}

// convertName 名前に変換する
func convertName(name string) Name {

	return Name{name}
}

// convertBirthday 誕生日に変換する
func convertBirthDay(birthDay string) BirthDay {

	d, _ := convertDateTime(birthDay)
	return BirthDay{value: d}
}

// validateName 名前の値を検証する
// 名前は下記の条件を満たす
//   - 値が入力されていること
func validateName(name string) base.ValidationResults {

	const itemName = "名前"
	r := base.NewValidationResults()

	if len(name) < 1 {

		r.AddRequiredError(itemName)
	}

	return r
}

// validateBirthDay 誕生日の値を検証する
// 誕生日は下記の条件を満たす
//   - 値が入力されていること
//   - YYYY-MM-DD形式であること
func validateBirthDay(birthDay string) base.ValidationResults {

	const itemName = "誕生日"
	r := base.NewValidationResults()

	if len(birthDay) < 1 {

		r.AddRequiredError(itemName)
		return r
	}

	_, br := convertDateTime(birthDay)
	if br != nil {

		r.AddDateFormatError(itemName)
	}

	return r
}

// validateExistUser 存在するユーザかを判定する
// 存在しない場合はエラーとして扱う
func validateExistUser(repository Repository, id ID) base.ValidationResults {

	r := base.NewValidationResults()

	_, er := repository.FindById(id)
	if er != nil {

		r.AddDataNotExistsError("ユーザ")
	}

	return r
}

// convertDateTime 日付型に変換する
func convertDateTime(value string) (time.Time, error) {

	loc, _ := time.LoadLocation("Asia/Tokyo")
	return time.ParseInLocation("2006-01-02", value, loc)
}
