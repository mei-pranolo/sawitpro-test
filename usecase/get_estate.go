package usecase

import (
	"context"

	m "github.com/SawitProRecruitment/UserService/types"
)

func (u *Usecase) GetEstateByID(ctx context.Context, id string) (estate m.Estate, err error) {
	return u.Repo.GetEstateByID(ctx, id)
}
