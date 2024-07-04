// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"

	m "github.com/SawitProRecruitment/UserService/types"
)

//go:generate mockgen --build_flags=--mod=mod -destination=interfaces.mock.gen.go -package=repository . RepositoryInterface
type RepositoryInterface interface {
	GetEstateByID(ctx context.Context, id string) (estate m.Estate, err error)
	CreateEstate(ctx context.Context, length int, width int) (id string, err error)
	CreateTree(ctx context.Context, estateID string, tree m.Tree) (id string, err error)
	GetTree(ctx context.Context, estateID string) (tree []m.Tree, err error)
}
