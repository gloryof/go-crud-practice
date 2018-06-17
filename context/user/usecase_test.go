package user

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUsecase_Register(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		param ModifyUserParam
	}
	tests := []struct {
		name       string
		repository Repository
		args       args
		want       ID
		wantErr    bool
	}{
		{
			name: "正常系",
			repository: createSaveMock(mockCtrl,
				User{
					name:     Name{"Junki"},
					birthDay: BirthDay{createDateTime("1986-12-16")},
				},
				createUserID(),
			),
			args: args{
				param: ModifyUserParam{Name: "Junki", BirthDay: "1986-12-16"},
			},
			want:    ID{value: 1000, numbered: true},
			wantErr: false,
		},
		{
			name: "入力チェックエラー",
			args: args{
				param: ModifyUserParam{Name: "", BirthDay: "1986-12-16"},
			},
			want:    ID{},
			wantErr: true,
		},
		{
			name:       "リポジトリエラー",
			repository: createSaveErrorMock(mockCtrl),
			args: args{
				param: ModifyUserParam{Name: "Yamada", BirthDay: "1986-12-16"},
			},
			want:    ID{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository: tt.repository,
			}
			got, err := u.Register(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_Update(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		id    uint64
		param ModifyUserParam
	}
	tests := []struct {
		name       string
		repository Repository
		args       args
		wantErr    bool
	}{
		{
			name: "正常系",
			repository: createUpdateMock(mockCtrl,
				User{
					id:       createUserID(),
					name:     Name{"Before"},
					birthDay: BirthDay{createDateTime("1982-01-01")},
				},
				User{
					id:       createUserID(),
					name:     Name{"Junki"},
					birthDay: BirthDay{createDateTime("1986-12-16")},
				},
			),
			args: args{
				id:    1000,
				param: ModifyUserParam{Name: "Junki", BirthDay: "1986-12-16"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				repository: tt.repository,
			}
			if err := u.Update(tt.args.id, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("Usecase.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecase_FindByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var idv uint64 = 1000
	eu := User{
		id:       ID{value: idv, numbered: true},
		name:     Name{value: "テスト"},
		birthDay: BirthDay{value: createDateTime("1986-12-16")},
	}

	type fields struct {
		repository Repository
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "正常系",
			fields: fields{
				repository: createFindMock(mockCtrl, idv, eu),
			},
			args:    args{id: idv},
			want:    eu,
			wantErr: false,
		},
		{
			name: "エラー系",
			fields: fields{
				repository: createFindErrorMock(mockCtrl, idv),
			},
			args:    args{id: idv},
			want:    User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository: tt.fields.repository,
			}
			got, err := u.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_DeleteByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		repository Repository
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			fields: fields{
				repository: createDeleteMock(mockCtrl, 1000),
			},
			args: args{
				id: 1000,
			},
			wantErr: false,
		},
		{
			name: "エラー系",
			fields: fields{
				repository: createDeleteMock(mockCtrl, 1000),
			},
			args: args{
				id: 1000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository: tt.fields.repository,
			}
			if err := u.DeleteByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Usecase.DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func createUserID() ID {
	return ID{value: 1000, numbered: true}
}

func createSaveMock(mockCtrl *gomock.Controller, nu User, id ID) *MockRepository {

	mr := NewMockRepository(mockCtrl)

	mr.EXPECT().Save(nu).Return(id, nil)

	return mr
}

func createSaveErrorMock(mockCtrl *gomock.Controller) *MockRepository {

	mr := NewMockRepository(mockCtrl)

	eu := User{
		name:     Name{value: "Yamada"},
		birthDay: BirthDay{value: createDateTime("1986-12-16")},
	}
	mr.EXPECT().Save(eu).Return(ID{}, errors.New("test"))

	return mr
}

func createUpdateMock(mockCtrl *gomock.Controller, base User, updated User) *MockRepository {

	mr := NewMockRepository(mockCtrl)

	mr.EXPECT().FindById(ID{value: base.id.GetValue()}).Return(base, nil)
	mr.EXPECT().Save(updated).Return(updated.id, nil)

	return mr
}

func createFindMock(mockCtrl *gomock.Controller, id uint64, base User) *MockRepository {

	mr := NewMockRepository(mockCtrl)

	mr.EXPECT().FindById(ID{value: id}).Return(base, nil)

	return mr
}

func createFindErrorMock(mockCtrl *gomock.Controller, id uint64) *MockRepository {

	mr := NewMockRepository(mockCtrl)

	mr.EXPECT().FindById(ID{value: id}).Return(User{}, errors.New("Test"))

	return mr
}

func createDeleteMock(mockCtrl *gomock.Controller, id uint64) *MockRepository {

	mr := NewMockRepository(mockCtrl)

	mr.EXPECT().DeleteByID(ID{value: id}).Return(nil)

	return mr
}

func createDeleteErrorMock(mockCtrl *gomock.Controller, id uint64) *MockRepository {

	mr := NewMockRepository(mockCtrl)

	mr.EXPECT().DeleteByID(ID{value: id}).Return(errors.New("Test"))

	return mr
}
