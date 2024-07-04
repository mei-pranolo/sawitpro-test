package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	repoMock "github.com/SawitProRecruitment/UserService/repository/mock"
	m "github.com/SawitProRecruitment/UserService/types"
	"go.uber.org/mock/gomock"
)

func TestUsecase_CreateTree(t *testing.T) {
	mockRepo := repoMock.NewMockRepositoryInterface(gomock.NewController(t))
	type args struct {
		ctx      context.Context
		estateID string
		tree     m.Tree
	}
	tests := []struct {
		name      string
		args      args
		wantId    string
		wantErr   bool
		repo      repository.RepositoryInterface
		mockCalls []func() *gomock.Call
	}{
		{
			name: "when all good, return id and no error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
				tree:     m.Tree{X: 1, Y: 1, Height: 2},
			},
			wantId:  "aaa",
			wantErr: false,
			repo:    mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{ID: "aaa", Length: 2, Width: 2}, nil)
				},
				func() *gomock.Call {
					return mockRepo.EXPECT().CreateTree(gomock.Any(), "aaa", m.Tree{X: 1, Y: 1, Height: 2}).Return("aaa", nil)
				},
			},
		},
		{
			name: "when estate not found, return error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
				tree:     m.Tree{X: 1, Y: 1, Height: 2},
			},
			wantId:  "",
			wantErr: true,
			repo:    mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{}, nil)
				},
			},
		},
		{
			name: "when get estate return error, return error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
				tree:     m.Tree{X: 1, Y: 1, Height: 2},
			},
			wantId:  "",
			wantErr: true,
			repo:    mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{}, errors.New("estate"))
				},
			},
		},
		{
			name: "when tree outside the estate frame, return error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
				tree:     m.Tree{X: 5, Y: 1, Height: 2},
			},
			wantId:  "",
			wantErr: true,
			repo:    mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{ID: "aaa", Length: 2, Width: 2}, nil)
				},
			},
		},
		{
			name: "when tree is too tall, return error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
				tree:     m.Tree{X: 1, Y: 1, Height: 32},
			},
			wantId:  "",
			wantErr: true,
			repo:    mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{ID: "aaa", Length: 2, Width: 2}, nil)
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

			gotId, err := u.CreateTree(tt.args.ctx, tt.args.estateID, tt.args.tree)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CreateTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Usecase.CreateTree() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
