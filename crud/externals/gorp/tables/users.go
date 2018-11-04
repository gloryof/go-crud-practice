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
	Name string `db:"name"`
	// BirthDay 誕生日
	BirthDay time.Time `db:"birthday"`
}

// UsersDao usersテーブルのDAO
type UsersDao interface {
	// SelectAll 全件取得
	SelectAll() []Users

	// SelectByID IDでSELECTする
	// データが取得できた場合はUsers型のデータが設定され、bool型はtrueで返る
	// データが取得できなかった場合はUsers型は初期値で、bool型はfalseで返る
	SelectByID(id uint64) (Users, bool)

	// Insert レコードの登録を行う
	// 登録に成功した場合はtrue、登録に失敗した場合はfalseを返す
	Insert(user Users) bool

	// Update レコードの更新を行う
	// 更新に成功した場合はtrue、更新に失敗した場合はfalseを返す
	Update(user Users) bool

	// DeleteByID IDをキーにレコードの削除を行う
	// 削除に成功した場合はtrue、削除に失敗した場合はfalseを返す
	DeleteByID(id uint64) bool

	// SelectNextID 次のIDを発行する
	// 発行に成功した場合はtrue、発行に失敗した場合はfalseを返す
	SelectNextID() (uint64, bool)
}

// UsersDaoGorpImpl usersテーブルのDAO
type UsersDaoGorpImpl struct {
	// DBMap DbMap
	DBMap *gorp.DbMap
}

// SelectAll UsersDaoの実装
func (dao UsersDaoGorpImpl) SelectAll() []Users {

	var result []Users
	dao.DBMap.Select(&result, "SELECT * FROM users")

	return result
}

// SelectByID UsersDaoの実装
func (dao UsersDaoGorpImpl) SelectByID(id uint64) (Users, bool) {

	r, err := dao.DBMap.Get(Users{}, id)

	if err != nil {

		return Users{}, false
	}

	ru := r.(*Users)
	return *ru, true
}

// Insert UsersDaoの実装
func (dao UsersDaoGorpImpl) Insert(user Users) bool {

	err := dao.DBMap.Insert(user)

	return err == nil
}

// Update UsersDaoの実装
func (dao UsersDaoGorpImpl) Update(user Users) bool {

	_, err := dao.DBMap.Update(user)

	return err == nil
}

// DeleteByID UsersDaoの実装
func (dao UsersDaoGorpImpl) DeleteByID(id uint64) bool {

	_, err := dao.DBMap.Exec("DELETE FROM users WHERE id = ?", id)

	return err == nil
}

// SelectNextID 次のIDを発行する
func (dao UsersDaoGorpImpl) SelectNextID() (uint64, bool) {

	i, error := dao.DBMap.SelectInt("SELECT NEXTVAL('user_id_seq')")

	if error != nil {

		return uint64(i), true
	}

	return 0, false
}
