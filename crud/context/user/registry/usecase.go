package registry

import "github.com/gloryof/go-crud-practice/crud/context/user/usecase"

// UsecaseResult registryの実行結果
type UsecaseResult struct {
	Usecase usecase.UserUsecase
}

// RegisterUsecase 依存性の登録を行う
func RegisterUsecase(infra InfraResult) UsecaseResult {
	return UsecaseResult{
		Usecase: usecase.New(infra.Repository),
	}
}
