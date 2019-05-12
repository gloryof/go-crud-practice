package echo

import (
	"github.com/gloryof/go-crud-practice/crud/config/registry"
	"github.com/labstack/echo/v4"
)

// routeUser ユーザに関するURLのルータ設定を行う
func routeUser(e *echo.Echo, handlers *registry.UserHandler) {

	ug := e.Group("/user")

	ug.GET("/list", handlers.UserList.ViewAll)
	ug.GET("/detail/:userID", handlers.UserDetail.ViewDetail)
	ug.GET("/detail/:userID/edit", handlers.UserDetail.ViewEdit)
	ug.POST("/detail/:userID/edit", handlers.UserDetail.ExecuteUpdating)
}
