package domain

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/gloryof/go-crud-practice/context/base"
	"github.com/golang/mock/gomock"
)

type inputConvertValues struct {
	name      string
	args      RegisterSpec
	wantUser  User
	wantError *base.ValidationResults
}

type inputValidateValues struct {
	name string
	args RegisterSpec
	want base.ValidationResults
}

func TestRegisterSpec_Validate(t *testing.T) {
	tests := createValidateValues()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, wantError %v", got, tt.want)
			}
		})
	}
}

func TestRegisterSpec_Convert(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	notNumberedID := ID{numbered: false, value: 0}

	tests := []struct {
		name  string
		args  RegisterSpec
		want  User
		want1 *base.ValidationResults
	}{

		{
			name: "正常パターン",
			args: RegisterSpec{"テスト", "2018-05-01"},
			want: User{
				id:       notNumberedID,
				name:     Name{value: "テスト"},
				birthDay: BirthDay{value: time.Date(2018, 05, 01, 0, 0, 0, 0, loc)},
			},
		},
		{
			name:  "入力エラー",
			args:  RegisterSpec{"", "2018-05-01"},
			want:  User{},
			want1: createExpectError1(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.args.Convert()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Convert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUpdateSpec_Validate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		ID         uint64
		Name       string
		BirthDay   string
		Repository Repository
	}
	tests := []struct {
		name   string
		fields fields
		want   base.ValidationResults
	}{
		{
			name: "正常系",
			fields: fields{
				ID:         1000,
				Name:       "テストユーザ",
				BirthDay:   "1986-12-16",
				Repository: createExistsUserMock(mockCtrl),
			},
			want: base.NewValidationResults(),
		},
		{
			name: "RegisterSpecの内容でエラー",
			fields: fields{
				ID:         1000,
				Name:       "",
				BirthDay:   "1986-12-16",
				Repository: createExistsUserMock(mockCtrl),
			},
			want: *createExpectError1(),
		},
		{
			name: "対象ユーザが存在しない",
			fields: fields{
				ID:         1000,
				Name:       "テストユーザ",
				BirthDay:   "1986-12-16",
				Repository: createNotExistsUserMock(mockCtrl, 1000),
			},
			want: *createExpectError4(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UpdateSpec{
				Name:       tt.fields.Name,
				BirthDay:   tt.fields.BirthDay,
				ID:         tt.fields.ID,
				Repository: tt.fields.Repository,
			}
			if got := s.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSpec.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateSpec_Convert(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		ID         uint64
		Name       string
		BirthDay   string
		Repository Repository
	}
	tests := []struct {
		name   string
		fields fields
		want   User
		want1  *base.ValidationResults
	}{
		{
			name: "正常系",
			fields: fields{
				ID:         1000,
				Name:       "テストユーザ",
				BirthDay:   "1986-12-16",
				Repository: createExistsUserMock(mockCtrl),
			},
			want: User{
				id:       ID{numbered: true, value: 1000},
				name:     Name{value: "テストユーザ"},
				birthDay: BirthDay{value: createDateTime("1986-12-16")},
			},
			want1: nil,
		},
		{
			name: "RegisterSpecの内容でエラー",
			fields: fields{
				ID:         1000,
				Name:       "",
				BirthDay:   "1986-12-16",
				Repository: createExistsUserMock(mockCtrl),
			},
			want:  User{},
			want1: createExpectError1(),
		},
		{
			name: "対象ユーザが存在しない",
			fields: fields{
				ID:         1000,
				Name:       "テストユーザ",
				BirthDay:   "1986-12-16",
				Repository: createNotExistsUserMock(mockCtrl, 1000),
			},
			want:  User{},
			want1: createExpectError4(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UpdateSpec{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				BirthDay:   tt.fields.BirthDay,
				Repository: tt.fields.Repository,
			}
			got, got1 := s.Convert()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSpec.Convert() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UpdateSpec.Convert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func createValidateValues() []inputValidateValues {

	return []inputValidateValues{
		{
			name: "正常パターン",
			args: RegisterSpec{"テスト", "2018-05-01"},
			want: base.NewValidationResults(),
		},
		{
			name: "正常パターン（閏年）",
			args: RegisterSpec{"テスト", "2016-02-29"},
			want: base.NewValidationResults(),
		},
		{
			name: "正常パターン（400年ごとの閏年）",
			args: RegisterSpec{"テスト", "2000-02-29"},
			want: base.NewValidationResults(),
		},
		{
			name: "名前が未入力",
			args: RegisterSpec{"", "2018-05-01"},
			want: *createExpectError1(),
		},
		{
			name: "誕生日が未入力",
			args: RegisterSpec{"テスト", ""},
			want: *createExpectError2(),
		},
		{
			name: "日付形式にアルファベットを含む",
			args: RegisterSpec{"テスト", "201a-05-20"},
			want: *createExpectError3(),
		},
		{
			name: "年数が4桁を超える",
			args: RegisterSpec{"テスト", "20018-05-20"},
			want: *createExpectError3(),
		},
		{
			name: "月が0",
			args: RegisterSpec{"テスト", "2018-00-20"},
			want: *createExpectError3(),
		},
		{
			name: "月が13",
			args: RegisterSpec{"テスト", "2018-13-20"},
			want: *createExpectError3(),
		},
		{
			name: "月の桁数が1",
			args: RegisterSpec{"テスト", "2018-1-20"},
			want: *createExpectError3(),
		},
		{
			name: "日が0",
			args: RegisterSpec{"テスト", "2018-05-00"},
			want: *createExpectError3(),
		},
		{
			name: "日が32",
			args: RegisterSpec{"テスト", "2018-05-32"},
			want: *createExpectError3(),
		},
		{
			name: "日が31（30日までの月）",
			args: RegisterSpec{"テスト", "2018-04-31"},
			want: *createExpectError3(),
		},
		{
			name: "日が29（閏年ではない2月）",
			args: RegisterSpec{"テスト", "2018-02-29"},
			want: *createExpectError3(),
		},
		{
			name: "日が29（閏年ではない2月）",
			args: RegisterSpec{"テスト", "2018-02-29"},
			want: *createExpectError3(),
		},
		{
			name: "日が1桁",
			args: RegisterSpec{"テスト", "2018-05-1"},
			want: *createExpectError3(),
		},
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
func createExpectError4() *base.ValidationResults {
	r := base.NewValidationResults()

	r.AddDataNotExistsError("ユーザ")

	return &r
}

func createExistsUserMock(mockCtrl *gomock.Controller) *MockRepository {

	sid := ID{value: 1000}

	base := User{
		id:       ID{value: sid.value, numbered: true},
		name:     Name{value: "変更前"},
		birthDay: BirthDay{value: createDateTime("1900-01-01")},
	}

	mr := NewMockRepository(mockCtrl)

	mr.EXPECT().FindById(sid).Return(base, nil)

	return mr
}

func createNotExistsUserMock(mockCtrl *gomock.Controller, id uint64) *MockRepository {

	mr := NewMockRepository(mockCtrl)

	uid := ID{value: id}

	mr.EXPECT().FindById(uid).Return(User{}, errors.New("検索エラー"))

	return mr
}
