package domain

import "errors"

type mockRepository struct {
	users []User
}

// FindByID Repositoryの内部モック実装
func (m mockRepository) FindByID(id ID) (User, error) {

	for _, u := range m.users {

		if u.GetID().GetValue() == id.GetValue() {

			return u, nil
		}
	}

	return User{}, errors.New("ユーザが見つかりませんでした")
}

// FindById Repositoryの内部モック実装
func (m mockRepository) Save(user User) (ID, error) {
	return user.GetID(), nil
}

// FindById Repositoryの内部モック実装
func (m mockRepository) DeleteByID(id ID) error {
	return nil
}
