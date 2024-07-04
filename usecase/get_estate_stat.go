package usecase

import (
	"context"
	"errors"
	"slices"

	m "github.com/SawitProRecruitment/UserService/types"
)

func countStat(trees []m.Tree) (stats m.Stats) {
	height := []int{}
	for _, t := range trees {
		height = append(height, t.Height)
	}
	slices.Sort(height)

	center_idx := len(height) / 2
	if len(height)%2 == 1 {
		stats.Median = height[center_idx]
	} else {
		stats.Median = (height[center_idx-1] + height[center_idx]) / 2
	}

	stats.Count = len(trees)
	stats.Max = slices.Max(height)
	stats.Min = slices.Min(height)
	return stats
}

func (u *Usecase) GetEstateStats(ctx context.Context, estateID string) (stat m.Stats, err error) {
	// check if estate exist
	estate, err := u.Repo.GetEstateByID(ctx, estateID)
	if err != nil {
		return
	}
	if estate.ID == "" {
		return m.Stats{}, errors.New("estate is not exist")
	}

	trees, err := u.Repo.GetTree(ctx, estateID)
	if err != nil {
		return
	}

	return countStat(trees), nil
}
