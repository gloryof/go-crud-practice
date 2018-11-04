package usecase

import "github.com/gloryof/go-crud-practice/crud/context/user/domain"

// ModifyUser ユーザを編集する
type ModifyUser struct {
	// repository ユーザリポジトリ
	repository domain.Repository
}

// NewModifyUser ユーザ編集ユースケースを作成する
func NewModifyUser(repository domain.Repository) ModifyUser {
	return ModifyUser{repository}
}

// ModifyUserParam Registerのパラメータ
type ModifyUserParam struct {
	Name     string
	BirthDay string
}

// Register 登録処理
// 登録に成功した場合はIDを返す。
// 登録に失敗した場合はerrorを返す
func (u ModifyUser) Register(param ModifyUserParam) (domain.ID, error) {

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
func (u ModifyUser) Update(id uint64, param ModifyUserParam) error {

	sp := domain.UpdateSpec{
		ID:         id,
		Name:       param.Name,
		BirthDay:   param.BirthDay,
		Repository: u.repository,
	}

	us, ve := sp.Convert()

	if ve != nil {

		return ve
	}

	_, re := u.repository.Save(us)

	return re
}

// DeleteByID ユーザIDに紐づくユーザを削除する
// 削除に失敗した場合はerrorを返す
func (u ModifyUser) DeleteByID(id uint64) error {

	return u.repository.DeleteByID(domain.NewID(id))
}
