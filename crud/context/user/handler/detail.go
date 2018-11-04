package handler

import (
	"net/http"
	"strconv"

	"github.com/gloryof/go-crud-practice/crud/context/user/usecase"
	"github.com/labstack/echo"
)

// UserDetail ユーザ詳細
type UserDetail struct {
	usecase *usecase.SearchUser
}

// NewUserDetail ユーザ一覧の作成
func NewUserDetail(usecase *usecase.SearchUser) UserDetail {

	return UserDetail{
		usecase: usecase,
	}
}

// UserDetailView ユーザ詳細表示情報
type UserDetailView struct {
	User UserInfo
}

// ViewDetail 詳細表示処理
func (u UserDetail) ViewDetail(c echo.Context) error {

	sid := c.Param("userID")
	uid, _ := strconv.ParseUint(sid, 10, 0)
	e, er := u.usecase.FindByID(uid)

	if er != nil {

		return er
	}

	view := UserDetailView{
		User: mapToUserInfo(e),
	}

	return c.Render(http.StatusOK, "user/detail", view)
}
