package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gloryof/go-crud-practice/crud/context/base"
	"github.com/gloryof/go-crud-practice/crud/context/user/usecase"
	"github.com/labstack/echo/v4"
)

// UserDetail ユーザ詳細
type UserDetail struct {
	viewUsecase *usecase.SearchUser
	editUsecase *usecase.ModifyUser
}

// NewUserDetail ユーザ一覧の作成
func NewUserDetail(viewUsecase *usecase.SearchUser, editUsecase *usecase.ModifyUser) UserDetail {

	return UserDetail{
		viewUsecase: viewUsecase,
		editUsecase: editUsecase,
	}
}

// UserDetailView ユーザ詳細表示情報
type UserDetailView struct {
	Errors []string
	User   UserInfo
}

// ViewDetail 詳細表示処理
func (u UserDetail) ViewDetail(c echo.Context) error {

	v, er := u.createView(c)

	if er != nil {

		return er
	}

	return c.Render(http.StatusOK, "user/detail", v)
}

// ViewEdit 編集表示処理
func (u UserDetail) ViewEdit(c echo.Context) error {

	v, er := u.createView(c)

	if er != nil {

		return er
	}

	return c.Render(http.StatusOK, "user/edit", v)
}

// ExecuteUpdating 更新処理
func (u UserDetail) ExecuteUpdating(c echo.Context) error {

	sid := c.Param("userID")
	uid, _ := strconv.ParseUint(sid, 10, 0)

	in := UserDetailView{
		User: UserInfo{
			ID:       uid,
			Name:     c.FormValue("name"),
			BirthDay: c.FormValue("birthDay"),
		},
	}

	md := usecase.ModifyUserParam{
		Name:     in.User.Name,
		BirthDay: in.User.BirthDay,
	}

	er := u.editUsecase.Update(uid, md)

	if er != nil {

		switch e := er.(type) {
		case *base.ValidationResults:
			bv := er.(*base.ValidationResults)
			in.Errors = bv.GetResults()
			return c.Render(http.StatusOK, "user/edit", in)
		default:
			return e
		}
	}

	return c.Redirect(http.StatusSeeOther, "/user/detail/"+fmt.Sprint(uid))
}

// createView ビューを作成する
func (u UserDetail) createView(c echo.Context) (UserDetailView, error) {

	sid := c.Param("userID")
	uid, _ := strconv.ParseUint(sid, 10, 0)
	e, er := u.viewUsecase.FindByID(uid)

	if er != nil {

		return UserDetailView{}, er
	}

	view := UserDetailView{
		User: mapToUserInfo(e),
	}

	return view, nil
}
