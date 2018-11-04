package usecase

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gloryof/go-crud-practice/crud/context/user/domain"
	"github.com/gloryof/go-crud-practice/tool/test/context/user/domain/mock"
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
		repository domain.Repository
		args       args
		want       domain.ID
		wantErr    bool
	}{
		{
			name: "正常系",
			repository: createSaveMock(mockCtrl,
				createTestNewUser("Junki", "1986-12-16"),
				domain.NewID(1000),
			),
			args: args{
				param: ModifyUserParam{Name: "Junki", BirthDay: "1986-12-16"},
			},
			want:    domain.NewID(1000),
			wantErr: false,
		},
		{
			name: "入力チェックエラー",
			args: args{
				param: ModifyUserParam{Name: "", BirthDay: "1986-12-16"},
			},
			want:    domain.ID{},
			wantErr: true,
		},
		{
			name:       "リポジトリエラー",
			repository: createSaveErrorMock(mockCtrl),
			args: args{
				param: ModifyUserParam{Name: "Yamada", BirthDay: "1986-12-16"},
			},
			want:    domain.ID{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := ModifyUser{
				repository: tt.repository,
			}
			got, err := u.Register(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.Register() = %v, want %v", got, tt.want)
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
		repository domain.Repository
		args       args
		wantErr    bool
	}{
		{
			name: "正常系",
			repository: createUpdateMock(mockCtrl,
				createTestExistUser(1000, "Before", "1982-01-01"),
				createTestExistUser(1000, "Junki", "1986-12-16"),
			),
			args: args{
				id:    1000,
				param: ModifyUserParam{Name: "Junki", BirthDay: "1986-12-16"},
			},
			wantErr: false,
		},
		{
			name: "入力チェックエラー",
			repository: createViewOnlyMock(mockCtrl,
				createTestExistUser(1000, "Before", "1982-01-01"),
			),
			args: args{
				id:    1000,
				param: ModifyUserParam{Name: "", BirthDay: "1986-12-16"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := ModifyUser{
				repository: tt.repository,
			}
			if err := u.Update(tt.args.id, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecase_DeleteByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		repository domain.Repository
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
			u := ModifyUser{
				repository: tt.fields.repository,
			}
			if err := u.DeleteByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func createSaveMock(mockCtrl *gomock.Controller, nu domain.User, id domain.ID) *domain_mock.MockRepository {

	mr := domain_mock.NewMockRepository(mockCtrl)

	mr.EXPECT().Save(nu).Return(id, nil)

	return mr
}

func createSaveErrorMock(mockCtrl *gomock.Controller) *domain_mock.MockRepository {

	mr := domain_mock.NewMockRepository(mockCtrl)

	eu := createTestNewUser("Yamada", "1986-12-16")

	mr.EXPECT().Save(eu).Return(domain.ID{}, errors.New("test"))

	return mr
}

func createUpdateMock(mockCtrl *gomock.Controller, base domain.User, updated domain.User) *domain_mock.MockRepository {

	mr := domain_mock.NewMockRepository(mockCtrl)
	id := domain.NewID(1000)

	mr.EXPECT().FindByID(id).Return(base, nil)
	mr.EXPECT().Save(updated).Return(updated.GetID(), nil)

	return mr
}
func createViewOnlyMock(mockCtrl *gomock.Controller, base domain.User) *domain_mock.MockRepository {

	mr := domain_mock.NewMockRepository(mockCtrl)
	id := domain.NewID(1000)

	mr.EXPECT().FindByID(id).Return(base, nil)

	return mr
}

func createTestNewUser(name string, birthDay string) domain.User {

	nv, _ := domain.NewName(name)
	bv, _ := domain.NewBirthDay(birthDay)

	return domain.NewUser(domain.ID{}, nv, bv)
}

func createTestExistUser(id uint64, name string, birthDay string) domain.User {

	nv, _ := domain.NewName(name)
	bv, _ := domain.NewBirthDay(birthDay)

	return domain.CreateExistUser(domain.NewID(id), nv, bv)
}

func createDeleteMock(mockCtrl *gomock.Controller, id uint64) *domain_mock.MockRepository {

	mr := domain_mock.NewMockRepository(mockCtrl)
	uid := domain.NewID(1000)

	mr.EXPECT().DeleteByID(uid).Return(nil)

	return mr
}

func createDeleteErrorMock(mockCtrl *gomock.Controller, id uint64) *domain_mock.MockRepository {

	mr := domain_mock.NewMockRepository(mockCtrl)
	uid := domain.NewID(1000)

	mr.EXPECT().DeleteByID(uid).Return(errors.New("Test"))

	return mr
}
