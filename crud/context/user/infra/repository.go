package infra

import (
	"errors"

	"github.com/gloryof/go-crud-practice/crud/context/user/domain"
	"github.com/gloryof/go-crud-practice/crud/externals/gorp/tables"
)

// RepositoryDBImpl RepositoryのDB実装
type RepositoryDBImpl struct {
	// userDao ユーザDAO
	userDao tables.UsersDao
}

// FindAll Repositoryの内部モック実装
func (r RepositoryDBImpl) FindAll() ([]domain.User, error) {

	records := r.userDao.SelectAll()
	results := []domain.User{}

	for _, record := range records {

		results = append(results, mapToEntity(record))
	}

	return results, nil
}

// FindByID Repositoryの内部モック実装
func (r RepositoryDBImpl) FindByID(id domain.ID) (domain.User, error) {

	record, ok := r.userDao.SelectByID(id.GetValue())

	if ok {

		return mapToEntity(record), nil
	}

	return domain.User{}, errors.New("ユーザが見つかりませんでした")
}

// Save Repositoryの内部モック実装
func (r RepositoryDBImpl) Save(user domain.User) (domain.ID, error) {

	if user.IsExists() {

		return r.update(user)
	}

	return r.insert(user)
}

// update 更新処理
func (r RepositoryDBImpl) update(user domain.User) (domain.ID, error) {

	id := user.GetID()
	record := mapToRecord(id.GetValue(), user)

	ok := r.userDao.Update(record)

	if ok {

		return id, nil
	}

	return domain.ID{}, errors.New("更新処理で失敗しました")
}

func (r RepositoryDBImpl) insert(user domain.User) (domain.ID, error) {

	id, ok := r.userDao.SelectNextID()

	if !ok {

		return domain.ID{}, errors.New("IDの採番に失敗しました。")
	}

	record := mapToRecord(id, user)

	r.userDao.Insert(record)

	if ok {

		return domain.NewID(id), nil
	}

	return domain.ID{}, errors.New("登録処理で失敗しました")
}

// DeleteByID Repositoryの内部モック実装
func (r RepositoryDBImpl) DeleteByID(id domain.ID) error {

	ok := r.userDao.DeleteByID(id.GetValue())

	if ok {

		return nil
	}

	return errors.New("削除処理に失敗しました")
}

// mapToEntity レコードからドメインエンティティに変換する
func mapToEntity(record tables.Users) domain.User {

	id := domain.NewID(record.ID)
	name, _ := domain.NewName(record.Name)
	birthDay := domain.ToBirthDay(record.BirthDay)

	return domain.CreateExistUser(id, name, birthDay)
}

// mapToRecord ドメインエンティティからレコードに変換する
func mapToRecord(id uint64, entity domain.User) tables.Users {

	return tables.Users{
		ID:       id,
		Name:     entity.GetName().GetValue(),
		BirthDay: entity.GetBirthDay().GetValue(),
	}
}
