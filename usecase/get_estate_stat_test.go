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

func TestUsecase_GetEstateStats(t *testing.T) {
	mockRepo := repoMock.NewMockRepositoryInterface(gomock.NewController(t))
	type args struct {
		ctx      context.Context
		estateID string
	}
	tests := []struct {
		name      string
		args      args
		wantStat  m.Stats
		wantErr   bool
		repo      repository.RepositoryInterface
		mockCalls []func() *gomock.Call
	}{
		{
			name: "when all good, return no error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
			},
			wantStat: m.Stats{Count: 4, Max: 5, Min: 3, Median: 4},
			wantErr:  false,
			repo:     mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{ID: "aaa", Length: 2, Width: 2}, nil)
				},
				func() *gomock.Call {
					return mockRepo.EXPECT().GetTree(gomock.Any(), "aaa").Return([]m.Tree{
						{X: 2, Y: 1, Height: 5},
						{X: 3, Y: 1, Height: 3},
						{X: 4, Y: 1, Height: 4},
						{X: 4, Y: 2, Height: 4},
					}, nil)
				},
			},
		},
		{
			name: "when get estate give error, return error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
			},
			wantStat: m.Stats{Count: 0, Max: 0, Min: 0, Median: 0},
			wantErr:  true,
			repo:     mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{ID: "aaa", Length: 2, Width: 2}, errors.New("estate"))
				},
			},
		},
		{
			name: "when get estate not found, return error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
			},
			wantStat: m.Stats{Count: 0, Max: 0, Min: 0, Median: 0},
			wantErr:  true,
			repo:     mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{}, nil)
				},
			},
		},
		{
			name: "when get tree give error, return error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
			},
			wantStat: m.Stats{Count: 0, Max: 0, Min: 0, Median: 0},
			wantErr:  true,
			repo:     mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{ID: "aaa", Length: 2, Width: 2}, nil)
				},
				func() *gomock.Call {
					return mockRepo.EXPECT().GetTree(gomock.Any(), "aaa").Return([]m.Tree{}, errors.New("tree"))
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
			gotStat, err := u.GetEstateStats(tt.args.ctx, tt.args.estateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetEstateStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStat, tt.wantStat) {
				t.Errorf("Usecase.GetEstateStats() = %v, want %v", gotStat, tt.wantStat)
			}
		})
	}
}
