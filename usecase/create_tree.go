package usecase

import (
	"context"
	"errors"

	m "github.com/SawitProRecruitment/UserService/types"
)

func (u *Usecase) CreateTree(ctx context.Context, estateID string, tree m.Tree) (id string, err error) {
	// safeguard
	estate, err := u.Repo.GetEstateByID(ctx, estateID)
	if err != nil {
		return
	}
	if estate.ID == "" {
		return "", errors.New("estate is not exist")
	}
	if tree.X > estate.Length || tree.Y > estate.Width {
		return "", errors.New("tree is outside estate")
	}

	return u.Repo.CreateTree(ctx, estateID, tree)
}
