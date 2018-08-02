package registry

import (
	"github.com/gloryof/go-crud-practice/crud/context/user/handler"
	"github.com/gloryof/go-crud-practice/crud/context/user/infra"
	"github.com/gloryof/go-crud-practice/crud/context/user/usecase"
	"github.com/gloryof/go-crud-practice/crud/externals/gorp/tables"
	"github.com/go-gorp/gorp"
)

// UserResult ユーザのregistry結果
type UserResult struct {
	// Infra インフラレイヤ
	Infra *UserInfra
	// Usecase ユースケースレイヤ
	Usecase *UserUsecase
	// Handler ハンドラレイヤ
	Handler *UserHandler
}

// UserInfra ユーザのインフラ群
type UserInfra struct {
	Repository *infra.RepositoryDBImpl
}

// UserHandler ユーザのハンドラ群
type UserHandler struct {
	// UserList ユーザ一覧
	UserList *handler.UserList
}

// UserUsecase ユーザのユースケース群
type UserUsecase struct {
	// Modify ユーザ編集
	Modify *usecase.ModifyUser
	// ユーザ検索
	Search *usecase.SearchUser
}

// registerUser ユーザの依存性の登録を行う
func registerUser(dbmap *gorp.DbMap) UserResult {

	i := registerUserInfra(dbmap)
	u := registerUserUsecase(i)
	h := registerUserHandler(u)

	return UserResult{
		Infra:   &i,
		Usecase: &u,
		Handler: &h,
	}
}

// registerUserHandler 依存性の登録を行う
func registerUserHandler(result UserUsecase) UserHandler {

	u := handler.NewUserList(result.Search)
	return UserHandler{
		UserList: &u,
	}
}

// registerUserInfra 依存性の登録を行う
func registerUserInfra(dbmap *gorp.DbMap) UserInfra {

	r := infra.NewRepositoryDBImpl(createUsersDao(dbmap))
	return UserInfra{
		Repository: &r,
	}
}

func createUsersDao(dbmap *gorp.DbMap) tables.UsersDaoGorpImpl {

	return tables.UsersDaoGorpImpl{
		DBMap: dbmap,
	}
}

// registerUserUsecase 依存性の登録を行う
func registerUserUsecase(infra UserInfra) UserUsecase {

	m := usecase.NewModifyUser(infra.Repository)
	s := usecase.NewSearchUser(infra.Repository)
	return UserUsecase{
		Modify: &m,
		Search: &s,
	}
}
