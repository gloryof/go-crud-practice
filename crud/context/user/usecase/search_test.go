package usecase

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gloryof/go-crud-practice/crud/context/user/domain"
	"github.com/gloryof/go-crud-practice/tool/test/context/user/domain/mock"
	"github.com/golang/mock/gomock"
)

func TestUsecase_FindByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var idv uint64 = 1000
	eu := createTestExistUser(1000, "テスト", "1986-12-16")

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
		want    domain.User
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
			want:    domain.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := SearchUser{
				repository: tt.fields.repository,
			}
			got, err := u.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createFindMock(mockCtrl *gomock.Controller, id uint64, base domain.User) *domain_mock.MockRepository {

	mr := domain_mock.NewMockRepository(mockCtrl)
	uid := domain.NewID(1000)

	mr.EXPECT().FindByID(uid).Return(base, nil)

	return mr
}

func createFindErrorMock(mockCtrl *gomock.Controller, id uint64) *domain_mock.MockRepository {

	mr := domain_mock.NewMockRepository(mockCtrl)
	uid := domain.NewID(1000)

	mr.EXPECT().FindByID(uid).Return(domain.User{}, errors.New("Test"))

	return mr
}
