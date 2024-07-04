package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	repoMock "github.com/SawitProRecruitment/UserService/repository/mock"
	m "github.com/SawitProRecruitment/UserService/types"
	"go.uber.org/mock/gomock"
)

func TestUsecase_GetEstateByID(t *testing.T) {
	mockRepo := repoMock.NewMockRepositoryInterface(gomock.NewController(t))
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name       string
		args       args
		wantEstate m.Estate
		wantErr    bool
		repo       repository.RepositoryInterface
		mockCalls  []func() *gomock.Call
	}{
		{
			name: "when all good, return no error",
			args: args{
				ctx: context.Background(),
				id:  "aaa",
			},
			wantEstate: m.Estate{ID: "aaa", Length: 2, Width: 2},
			wantErr:    false,
			repo:       mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{ID: "aaa", Length: 2, Width: 2}, nil)
				},
			},
		},
		{
			name: "when repo give error, return error",
			args: args{
				ctx: context.Background(),
				id:  "aaa",
			},
			wantEstate: m.Estate{},
			wantErr:    true,
			repo:       mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{}, errors.New("estate"))
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockCalls != nil {
				for _, call := range tt.mockCalls {
					call()
				}
			}

			u := &Usecase{
				Repo: tt.repo,
			}

			gotEstate, err := u.GetEstateByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetEstateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEstate, tt.wantEstate) {
				t.Errorf("Usecase.GetEstateByID() = %v, want %v", gotEstate, tt.wantEstate)
			}
		})
	}
}
