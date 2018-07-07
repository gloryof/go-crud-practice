package router

import (
	"github.com/gloryof/go-crud-practice/crud/context/user/registry"
	"github.com/gloryof/go-crud-practice/crud/externals"
	"github.com/labstack/echo"
)

// User ユーザに関するURLのルータ設定を行う
func User(c externals.Context, g *echo.Group) {

	infra := registry.RegisterInfra(c.DB.DBMap)
	usecase := registry.RegisterUsecase(infra)
	handler := registry.RegisterHandler(usecase)

	g.GET("/list", handler.UserList.ViewAll)
}
