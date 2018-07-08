package handler

import (
	"github.com/gloryof/go-crud-practice/crud/context/user/domain"
)

// UserInfo ユーザ情報
type UserInfo struct {

	// ID ID
	ID uint64

	// Name 名前
	Name string

	// BirthDay 誕生日
	BirthDay string
}

// mapToUserInfo ユーザエンティティをユーザ情報に変換する
func mapToUserInfo(user domain.User) UserInfo {
	return UserInfo{
		ID:       user.GetID().GetValue(),
		Name:     user.GetName().GetValue(),
		BirthDay: user.GetBirthDay().GetValue().Format("2006-01-02"),
	}
}
