package usecase

import "github.com/SawitProRecruitment/UserService/repository"

type Usecase struct {
	Repo repository.RepositoryInterface
}

func NewUsecase(repo repository.RepositoryInterface) *Usecase {
	return &Usecase{Repo: repo}
}
