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

func Test_countTraveledDistance(t *testing.T) {
	type args struct {
		trees     []m.Tree
		maxLength int
		maxWidth  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "all good",
			args: args{
				trees: []m.Tree{
					{X: 2, Y: 1, Height: 5},
					{X: 3, Y: 1, Height: 3},
					{X: 4, Y: 1, Height: 4},
					{X: 4, Y: 2, Height: 4},
				},
				maxLength: 5,
				maxWidth:  2,
			},
			want: 104,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTraveledDistance(tt.args.trees, tt.args.maxLength, tt.args.maxWidth); got != tt.want {
				t.Errorf("countTraveledDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetDroneDistance(t *testing.T) {
	mockRepo := repoMock.NewMockRepositoryInterface(gomock.NewController(t))
	type args struct {
		ctx      context.Context
		estateID string
	}
	tests := []struct {
		name         string
		args         args
		wantDistance int
		wantErr      bool
		repo         repository.RepositoryInterface
		mockCalls    []func() *gomock.Call
	}{
		{
			name: "when all good, return no error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
			},
			wantDistance: 54,
			wantErr:      false,
			repo:         mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{ID: "aaa", Length: 5, Width: 1}, nil)
				},
				func() *gomock.Call {
					return mockRepo.EXPECT().GetTree(gomock.Any(), "aaa").Return([]m.Tree{
						{X: 2, Y: 1, Height: 5},
						{X: 3, Y: 1, Height: 3},
						{X: 4, Y: 1, Height: 4}}, nil)
				},
			},
		},
		{
			name: "when estate not found, return error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
			},
			wantDistance: 0,
			wantErr:      true,
			repo:         mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{}, nil)
				},
			},
		},
		{
			name: "when get estate give error, return error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
			},
			wantDistance: 0,
			wantErr:      true,
			repo:         mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{}, errors.New("estate"))
				},
			},
		},
		{
			name: "when get tree give error, return error",
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
			},
			wantDistance: 0,
			wantErr:      true,
			repo:         mockRepo,
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return mockRepo.EXPECT().GetEstateByID(gomock.Any(), "aaa").Return(m.Estate{ID: "aaa", Length: 5, Width: 1}, nil)
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

			gotDistance, err := u.GetDroneDistance(tt.args.ctx, tt.args.estateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetDroneDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDistance != tt.wantDistance {
				t.Errorf("Usecase.GetDroneDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
