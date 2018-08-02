package handler

import (
	"net/http"

	"github.com/gloryof/go-crud-practice/crud/context/user/usecase"
	"github.com/labstack/echo"
)

// UserList ユーザ一覧
type UserList struct {
	usecase *usecase.SearchUser
}

// NewUserList ユーザ一覧の作成
func NewUserList(usecase *usecase.SearchUser) UserList {

	return UserList{
		usecase: usecase,
	}
}

// UserListView ユーザ一覧表示情報
type UserListView struct {
	Users []UserInfo
}

// ViewAll 全件表示処理
func (u UserList) ViewAll(c echo.Context) error {

	es, er := u.usecase.FindAll()

	if er != nil {

		return er
	}

	us := []UserInfo{}
	for _, e := range es {

		us = append(us, mapToUserInfo(e))
	}

	view := UserListView{
		Users: us,
	}

	return c.Render(http.StatusOK, "user/list", view)
}
