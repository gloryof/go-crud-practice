package user

// Usecase ユーザに関するユースケース
type Usecase struct {
	// repository ユーザリポジトリ
	repository Repository
}

// New ユースケースを作成する
func New(repository Repository) Usecase {
	return Usecase{repository}
}
