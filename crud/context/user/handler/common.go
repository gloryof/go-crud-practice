package handler

import (
	"time"

	"github.com/gloryof/go-crud-practice/crud/context/user/domain"
)

// UserInfo ユーザ情報
type UserInfo struct {

	// ID ID
	ID uint64

	// Name 名前
	Name string

	// BirthDay 誕生日
	BirthDay time.Time
}

// mapToUserInfo ユーザエンティティをユーザ情報に変換する
func mapToUserInfo(user domain.User) UserInfo {
	return UserInfo{
		ID:       user.GetID().GetValue(),
		Name:     user.GetName().GetValue(),
		BirthDay: user.GetBirthDay().GetValue(),
	}
}
