package tables

import (
	"time"

	"github.com/go-gorp/gorp"
)

// Users usersテーブルの構造体
type Users struct {
	// ID ユーザのID
	ID uint64 `db:"id, primarykey"`
	// Name 名前
	Name string
	// BirthDay 誕生日
	BirthDay time.Time
}

// UsersDao usersテーブルのDAO
type UsersDao struct {
	// DBMap DbMap
	DBMap *gorp.DbMap
}

// SelectAll 全件取得
func (dao UsersDao) SelectAll() []Users {

	var result []Users
	dao.DBMap.Select(&result, "SELECT * FROM users")

	return result
}

// SelectByID IDでSELECTする
// データが取得できた場合はUsers型のデータが設定され、bool型はtrueで返る
// データが取得できなかった場合はUsers型は初期値で、bool型はfalseで返る
func (dao UsersDao) SelectByID(id uint64) (Users, bool) {

	r, err := dao.DBMap.Get(Users{}, id)

	if err != nil {

		return r.(Users), false
	}

	return Users{}, true
}

// Insert レコードの登録を行う
// 登録に成功した場合はtrue、登録に失敗した場合はfalseを返す
func (dao UsersDao) Insert(user Users) bool {

	err := dao.DBMap.Insert(user)

	return err == nil
}

// Update レコードの更新を行う
// 更新に成功した場合はtrue、更新に失敗した場合はfalseを返す
func (dao UsersDao) Update(user Users) bool {

	_, err := dao.DBMap.Update(user)

	return err == nil
}

// DeleteByID IDをキーにレコードの削除を行う
// 削除に成功した場合はtrue、削除に失敗した場合はfalseを返す
func (dao UsersDao) DeleteByID(id uint64) bool {

	_, err := dao.DBMap.Exec("DELETE FROM users WHERE id = ?", id)

	return err == nil
}

// SelectNextID 次のIDを発行する
func (dao UsersDao) SelectNextID() (uint64, bool) {

	i, error := dao.DBMap.SelectInt("SELECT NEXTVAL('user_id_seq')")

	if error != nil {

		return uint64(i), true
	}

	return 0, false
}
