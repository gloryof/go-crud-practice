package user

import (
	"reflect"
	"testing"
	"time"

	"github.com/gloryof/go-crud-practice/context/base"
)

func TestExistsUser(t *testing.T) {

	var idValue uint64 = 1234
	expectedID := ID{numbered: true, value: idValue}
	expectedName := Name{value: "Test1"}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	expectedBirthDay := BirthDay{value: time.Date(2018, 05, 01, 0, 0, 0, 0, loc)}

	sut := ExistsUser(idValue, expectedName, expectedBirthDay)

	if !reflect.DeepEqual(sut.GetID(), expectedID) {
		t.Errorf("ExistsUser() got = %v, want %v", sut.GetID(), expectedID)
	}

	if !reflect.DeepEqual(sut.GetName(), expectedName) {
		t.Errorf("ExistsUser() got = %v, want %v", sut.GetName(), expectedName)
	}

	if !reflect.DeepEqual(sut.GetBirthDay(), expectedBirthDay) {
		t.Errorf("ExistsUser() got = %v, want %v", sut.GetBirthDay(), expectedBirthDay)
	}
}

func TestNewUser(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	notNumberedID := ID{numbered: false, value: 0}
	type args struct {
		name     string
		birthDay string
	}
	tests := []struct {
		name  string
		args  args
		want  User
		want1 *base.ValidationResults
	}{
		{
			name: "正常パターン",
			args: args{"テスト", "2018-05-01"},
			want: User{
				id:       notNumberedID,
				name:     Name{value: "テスト"},
				birthDay: BirthDay{value: time.Date(2018, 05, 01, 0, 0, 0, 0, loc)},
			},
		},
		{
			name: "正常パターン（閏年）",
			args: args{"テスト", "2016-02-29"},
			want: User{
				id:       notNumberedID,
				name:     Name{value: "テスト"},
				birthDay: BirthDay{value: time.Date(2016, 02, 29, 0, 0, 0, 0, loc)},
			},
		},
		{
			name: "正常パターン（400年ごとの閏年）",
			args: args{"テスト", "2000-02-29"},
			want: User{
				id:       notNumberedID,
				name:     Name{value: "テスト"},
				birthDay: BirthDay{value: time.Date(2000, 02, 29, 0, 0, 0, 0, loc)},
			},
		},
		{
			name:  "名前が未入力",
			args:  args{"", "2018-05-01"},
			want:  User{},
			want1: createExpectError1(),
		},
		{
			name:  "誕生日が未入力",
			args:  args{"テスト", ""},
			want:  User{},
			want1: createExpectError2(),
		},
		{
			name:  "日付形式にアルファベットを含む",
			args:  args{"テスト", "201a-05-20"},
			want:  User{},
			want1: createExpectError3(),
		},
		{
			name:  "年数が4桁を超える",
			args:  args{"テスト", "20018-05-20"},
			want:  User{},
			want1: createExpectError3(),
		},
		{
			name:  "月が0",
			args:  args{"テスト", "2018-00-20"},
			want:  User{},
			want1: createExpectError3(),
		},
		{
			name:  "月が13",
			args:  args{"テスト", "2018-13-20"},
			want:  User{},
			want1: createExpectError3(),
		},
		{
			name:  "月の桁数が1",
			args:  args{"テスト", "2018-1-20"},
			want:  User{},
			want1: createExpectError3(),
		},
		{
			name:  "日が0",
			args:  args{"テスト", "2018-05-00"},
			want:  User{},
			want1: createExpectError3(),
		},
		{
			name:  "日が32",
			args:  args{"テスト", "2018-05-32"},
			want:  User{},
			want1: createExpectError3(),
		},
		{
			name:  "日が31（30日までの月）",
			args:  args{"テスト", "2018-04-31"},
			want:  User{},
			want1: createExpectError3(),
		},
		{
			name:  "日が29（閏年ではない2月）",
			args:  args{"テスト", "2018-02-29"},
			want:  User{},
			want1: createExpectError3(),
		},
		{
			name:  "日が29（閏年ではない2月）",
			args:  args{"テスト", "2018-02-29"},
			want:  User{},
			want1: createExpectError3(),
		},
		{
			name:  "日が1桁",
			args:  args{"テスト", "2018-05-1"},
			want:  User{},
			want1: createExpectError3(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewUser(tt.args.name, tt.args.birthDay)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func createExpectError1() *base.ValidationResults {
	r := base.NewValidationResults()

	r.AddRequiredError("名前")

	return &r
}

func createExpectError2() *base.ValidationResults {
	r := base.NewValidationResults()

	r.AddRequiredError("誕生日")

	return &r
}

func createExpectError3() *base.ValidationResults {
	r := base.NewValidationResults()

	r.AddDateFormatError("誕生日")

	return &r
}
