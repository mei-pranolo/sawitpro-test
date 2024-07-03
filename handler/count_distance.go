package handler

import m "github.com/mei-pranolo/sawitpro-test/repository"

func CountTraveledDistance(trees []m.Tree, maxLength int, maxWidth int) int {
	total := 0
	y := 1
	x := 1

	curr_height := 0

	for y <= maxWidth {
		for x <= maxLength {
			for _, t := range trees {
				if t.x == x && t.y == y {
					if curr_height > t.height+1 {
						going_down := curr_height - (t.height + 1)
						total = total + going_down
						curr_height = curr_height - going_down
					} else if curr_height < t.height+1 {
						going_up := (t.height + 1) - curr_height
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
