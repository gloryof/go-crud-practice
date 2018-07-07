package registry

import "github.com/gloryof/go-crud-practice/crud/context/user/handler"

// HandlerResult registryの実行結果
type HandlerResult struct {
	// UserList ユーザ一覧
	UserList handler.UserList
}

// RegisterHandler 依存性の登録を行う
func RegisterHandler(result UsecaseResult) HandlerResult {

	return HandlerResult{
		UserList: handler.NewUserList(result.Usecase),
	}
}
