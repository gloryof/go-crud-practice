package domain

// Repository ユーザリポジトリ
type Repository interface {
	// FindAll IDをキーにユーザを探す
	// ユーザが存在する場合はUserを返す
	// 取得処理に失敗した場合はエラーを返す
	FindAll() ([]User, error)

	// FindByID IDをキーにユーザを探す
	// ユーザが存在する場合はUserを返す
	// ユーザが存在しない場合はエラーを返す
	FindByID(id ID) (User, error)

	// Save ユーザの保存を行う
	// 保存されたIDを返す
	// 保存処理に失敗した場合はエラーを返す
	Save(user User) (ID, error)

	// DeleteByID ユーザの削除を行う
	// 削除処理に失敗した場合はエラーを返す
	DeleteByID(id ID) error
}
