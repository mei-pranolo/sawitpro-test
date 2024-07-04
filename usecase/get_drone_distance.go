package usecase

import (
	"context"
	"errors"

	m "github.com/SawitProRecruitment/UserService/types"
)

func countTraveledDistance(trees []m.Tree, maxLength int, maxWidth int) int {
	total := 0
	y := 1
	x := 1

	curr_height := 0

	for y <= maxWidth {
		for x <= maxLength {
			for _, t := range trees {
				if t.X == x && t.Y == y {
					if curr_height > t.Height+1 {
						going_down := curr_height - (t.Height + 1)
						total = total + going_down
						curr_height = curr_height - going_down
					} else if curr_height < t.Height+1 {
						going_up := (t.Height + 1) - curr_height
						curr_height = curr_height + going_up
						total = total + going_up
					}
					break
				}
			}

			if x != maxLength {
				total = total + 10
			}
			x++
		}
		if y != maxWidth {
			total = total + 10
		} else {
			total = total + curr_height
		}
		x = 1
		y++
	}

	return total
}

func (u *Usecase) GetDroneDistance(ctx context.Context, estateID string) (distance int, err error) {
	// check if estate exist
	estate, err := u.Repo.GetEstateByID(ctx, estateID)
	if err != nil {
		return
	}
	if estate.ID == "" {
		return 0, errors.New("estate is not exist")
	}

	trees, err := u.Repo.GetTree(ctx, estateID)
	if err != nil {
		return
	}
	return countTraveledDistance(trees, estate.Length, estate.Width), nil
}
