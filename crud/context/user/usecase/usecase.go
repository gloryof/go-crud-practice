package user

import (
	"github.com/gloryof/go-crud-practice/crud/context/user/domain"
)

// Usecase ユーザに関するユースケース
type Usecase struct {
	// repository ユーザリポジトリ
	repository domain.Repository
}

// New ユースケースを作成する
func New(repository domain.Repository) Usecase {
	return Usecase{repository}
}

// ModifyUserParam Registerのパラメータ
type ModifyUserParam struct {
	Name     string
	BirthDay string
}

// Register 登録処理
// 登録に成功した場合はIDを返す。
// 登録に失敗した場合はerrorを返す
func (u Usecase) Register(param ModifyUserParam) (domain.ID, error) {

	sp := domain.RegisterSpec{
		Name:     param.Name,
		BirthDay: param.BirthDay,
	}
	us, ve := sp.Convert()

	if ve != nil {

		return domain.ID{}, ve
	}

	return u.repository.Save(us)
}

// Update 更新処理
// 更新に失敗した場合はerrorを返す
func (u Usecase) Update(id uint64, param ModifyUserParam) error {

	sp := domain.UpdateSpec{
		ID:         id,
		Name:       param.Name,
		BirthDay:   param.BirthDay,
		Repository: u.repository,
	}

	us, _ := sp.Convert()

	_, re := u.repository.Save(us)

	return re
}

// FindByID ユーザIDでユーザを検索する
// ユーザが存在しない場合はエラーを返す
func (u Usecase) FindByID(id uint64) (domain.User, error) {

	return u.repository.FindByID(domain.NewID(id))
}

// DeleteByID ユーザIDに紐づくユーザを削除する
// 削除に失敗した場合はerrorを返す
func (u Usecase) DeleteByID(id uint64) error {

	return u.repository.DeleteByID(domain.NewID(id))
}
