package usecase

import (
	"context"

	m "github.com/SawitProRecruitment/UserService/types"
)

//go:generate mockgen --build_flags=--mod=mod -destination=interfaces.mock.gen.go -package=usecase . UsecaseInterface
type UsecaseInterface interface {
	GetEstateByID(ctx context.Context, id string) (estate m.Estate, err error)
	CreateEstate(ctx context.Context, length int, width int) (id string, err error)
	CreateTree(ctx context.Context, estateID string, tree m.Tree) (id string, err error)

	GetEstateStats(ctx context.Context, estateID string) (stat m.Stats, err error)
	GetDroneDistance(ctx context.Context, estateID string) (distance int, err error)
}
