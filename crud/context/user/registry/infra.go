package registry

import (
	"github.com/gloryof/go-crud-practice/crud/context/user/infra"
	"github.com/gloryof/go-crud-practice/crud/externals/gorp/tables"
	"github.com/go-gorp/gorp"
)

// InfraResult registryの実行結果
type InfraResult struct {
	Repository infra.RepositoryDBImpl
}

// RegisterInfra 依存性の登録を行う
func RegisterInfra(dbmap *gorp.DbMap) InfraResult {

	return InfraResult{
		Repository: infra.NewRepositoryDBImpl(createUsersDao(dbmap)),
	}
}

func createUsersDao(dbmap *gorp.DbMap) tables.UsersDaoGorpImpl {

	return tables.UsersDaoGorpImpl{
		DBMap: dbmap,
	}
}
