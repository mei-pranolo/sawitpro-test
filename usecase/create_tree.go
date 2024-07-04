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
	if tree.Height > 30 || tree.Height < 1 {
		return "", errors.New("tree's height is not in range")
	}
	if tree.X > estate.Length || tree.X < 1 ||
		tree.Y > estate.Width || tree.Y < 1 {
		return "", errors.New("tree is outside estate")
	}

	return u.Repo.CreateTree(ctx, estateID, tree)
}
