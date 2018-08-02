package usecase

import (
	"github.com/gloryof/go-crud-practice/crud/context/user/domain"
)

// SearchUser ユーザを検索する
type SearchUser struct {
	// repository ユーザリポジトリ
	repository domain.Repository
}

// NewSearchUser ユーザ検索ユースケースを作成する
func NewSearchUser(repository domain.Repository) SearchUser {

	return SearchUser{repository}
}

// FindByID ユーザIDでユーザを検索する
// ユーザが存在しない場合はエラーを返す
func (u SearchUser) FindByID(id uint64) (domain.User, error) {

	return u.repository.FindByID(domain.NewID(id))
}

// FindAll 全てのユーザを取得する
func (u SearchUser) FindAll() ([]domain.User, error) {

	return u.repository.FindAll()
}
