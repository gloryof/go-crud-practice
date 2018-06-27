package domain

import (
	"reflect"
	"testing"

	"github.com/gloryof/go-crud-practice/crud/context/base"
)

func TestRegisterSpec_Validate(t *testing.T) {
	type fields struct {
		Name     string
		BirthDay string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "正常系",
			fields: fields{
				Name:     "テスト",
				BirthDay: "1986-12-16",
			},
			want: false,
		},
		{
			name: "名前にエラーがある場合",
			fields: fields{
				Name:     "",
				BirthDay: "1986-12-16",
			},
			want: true,
		},
		{
			name: "誕生日にエラーがある場合",
			fields: fields{
				Name:     "テスト",
				BirthDay: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RegisterSpec{
				Name:     tt.fields.Name,
				BirthDay: tt.fields.BirthDay,
			}
			if got := s.Validate(); !reflect.DeepEqual(got.HasError(), tt.want) {
				t.Errorf("RegisterSpec.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterSpec_Convert(t *testing.T) {
	uv, _ := base.ConvertDateTime("1986-12-16")

	type fields struct {
		Name     string
		BirthDay string
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
				Name:     "テスト",
				BirthDay: "1986-12-16",
			},
			want: User{
				id:         ID{},
				name:       Name{value: "テスト"},
				birthDay:   BirthDay{value: uv},
				registered: false,
			},
			want1: nil,
		},
		{
			name: "エラーがある場合",
			fields: fields{
				Name:     "",
				BirthDay: "1986-12-16",
			},
			want:  User{},
			want1: createConvertError(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RegisterSpec{
				Name:     tt.fields.Name,
				BirthDay: tt.fields.BirthDay,
			}
			got, got1 := s.Convert()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterSpec.Convert() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegisterSpec.Convert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUpdateSpec_Validate(t *testing.T) {

	t.Run("個別の入力値に対するエラー", func(t *testing.T) {
		bv, _ := base.ConvertDateTime("1900-01-01")
		mr := mockRepository{
			users: []User{
				User{
					id:         ID{value: 1000},
					name:       Name{value: "既存データ"},
					birthDay:   BirthDay{value: bv},
					registered: true,
				},
			},
		}

		type fields struct {
			ID         uint64
			Name       string
			BirthDay   string
			Repository Repository
		}
		tests := []struct {
			name   string
			fields fields
			want   bool
		}{
			{
				name: "正常系",
				fields: fields{
					ID:         1000,
					Name:       "テスト",
					BirthDay:   "1986-12-16",
					Repository: mr,
				},
				want: false,
			},
			{
				name: "名前にエラーがある場合",
				fields: fields{
					ID:         1000,
					Name:       "",
					BirthDay:   "1986-12-16",
					Repository: mr,
				},
				want: true,
			},
			{
				name: "誕生日にエラーがある場合",
				fields: fields{
					ID:         1000,
					Name:       "テスト",
					BirthDay:   "",
					Repository: mr,
				},
				want: true,
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
				if got := s.Validate(); !reflect.DeepEqual(got.HasError(), tt.want) {
					t.Errorf("UpdateSpec.Validate() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("個別の入力値以外に対するエラー", func(t *testing.T) {
		bv, _ := base.ConvertDateTime("1900-01-01")
		mr := mockRepository{
			users: []User{
				User{
					id:         ID{value: 1000},
					name:       Name{value: "既存データ"},
					birthDay:   BirthDay{value: bv},
					registered: true,
				},
			},
		}

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
				name: "既存データに存在する",
				fields: fields{
					ID:         1000,
					Name:       "テスト",
					BirthDay:   "1986-12-16",
					Repository: mr,
				},
				want: base.NewValidationResults(),
			},
			{
				name: "存在しないユーザの場合",
				fields: fields{
					ID:         2000,
					Name:       "テスト",
					BirthDay:   "1986-12-16",
					Repository: mr,
				},
				want: createNotExistsUser(),
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
				if got := s.Validate(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("UpdateSpec.Validate() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
func TestUpdateSpec_Convert(t *testing.T) {
	bv, _ := base.ConvertDateTime("1900-01-01")
	mr := mockRepository{
		users: []User{
			User{
				id:         ID{value: 1000},
				name:       Name{value: "既存データ"},
				birthDay:   BirthDay{value: bv},
				registered: true,
			},
		},
	}

	uv, _ := base.ConvertDateTime("1986-12-16")
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
				Name:       "テスト",
				BirthDay:   "1986-12-16",
				Repository: mr,
			},
			want: User{
				id:         ID{value: 1000},
				name:       Name{value: "テスト"},
				birthDay:   BirthDay{value: uv},
				registered: true,
			},
			want1: nil,
		},
		{
			name: "エラーがある場合",
			fields: fields{
				ID:         1000,
				Name:       "",
				BirthDay:   "1986-12-16",
				Repository: mr,
			},
			want:  User{},
			want1: createConvertError(),
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

func createNotExistsUser() base.ValidationResults {

	r := base.NewValidationResults()

	r.AddDataNotExistsError("ユーザ")

	return r
}

func createConvertError() *base.ValidationResults {

	r := base.NewValidationResults()

	r.AddRequiredError("名前")

	return &r
}
