package usecase

import (
	"context"
	"errors"
)

func (u *Usecase) CreateEstate(ctx context.Context, length int, width int) (id string, err error) {
	if length < 1 || length >= 50000 {
		return "", errors.New("length limit exceeded")
	}
	if width < 1 || width >= 50000 {
		return "", errors.New("width limit exceeded")
	}
	return u.Repo.CreateEstate(ctx, length, width)
}
