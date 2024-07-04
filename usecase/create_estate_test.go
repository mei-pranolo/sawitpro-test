package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	repoMock "github.com/SawitProRecruitment/UserService/repository/mock"
	"go.uber.org/mock/gomock"
)

func TestUsecase_CreateEstate(t *testing.T) {
	mockRepo := repoMock.NewMockRepositoryInterface(gomock.NewController(t))
	type args struct {
		ctx    context.Context
		length int
		width  int
	}
	tests := []struct {
		name      string
		repo      repository.RepositoryInterface
		args      args
		wantId    string
		wantErr   bool
		mockCalls []func() *gomock.Call
	}{
		{
			name: "when all good, return no error and id",
			repo: mockRepo,
			args: args{
				ctx:    context.Background(),
				length: 2,
				width:  2,
			},
			wantId:  "aabb",
			wantErr: false,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().CreateEstate(gomock.Any(), 2, 2).Return("aabb", nil)
				},
			},
		},
		{
			name: "when repo give error, return error",
			repo: mockRepo,
			args: args{
				ctx:    context.Background(),
				length: 2,
				width:  2,
			},
			wantId:  "",
			wantErr: true,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().CreateEstate(gomock.Any(), 2, 2).Return("", errors.New("db"))
				},
			},
		},
		{
			name: "when length not suitable, return error",
			repo: mockRepo,
			args: args{
				ctx:    context.Background(),
				length: -2,
				width:  2,
			},
			wantId:    "",
			wantErr:   true,
			mockCalls: []func() *gomock.Call{},
		},
		{
			name: "when width not suitable, return error",
			repo: mockRepo,
			args: args{
				ctx:    context.Background(),
				length: 2,
				width:  -2,
			},
			wantId:    "",
			wantErr:   true,
			mockCalls: []func() *gomock.Call{},
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

			gotId, err := u.CreateEstate(tt.args.ctx, tt.args.length, tt.args.width)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CreateEstate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Usecase.CreateEstate() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
