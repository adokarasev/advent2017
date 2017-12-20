package defragmentation

import (
	"strconv"

	"github.com/adokarasev/advent2017/knothash"
)

// CountUsedSquares returns the number of user squares
func CountUsedSquares(in []string) int64 {
	var totalUsed int64
	for _, s := range in {
		hash := knothash.Hash(256, s)
		for _, r := range hash {
			for n, _ := strconv.ParseInt(string(r), 16, 8); n != 0; n = n >> 1 {
				totalUsed += n & 1
			}
		}
	}
	return totalUsed
}

// CountGroup calculates a number of groups in the grid
func CountGroup(in []string) int {
	grid := make(map[int]int)
	var g int
	for _, s := range in {
		hash := knothash.Hash(256, s)
		for _, r := range hash {
			n, _ := strconv.ParseInt(string(r), 16, 8)
			for t := 0; t < 4; t++ {
				if n&1 == 1 {
					grid = addGroup(grid, g)
				}
				n = n >> 1
				g++
			}
		}
	}

	groups := make(map[int]bool)
	for _, g := range grid {
		if _, ok := groups[g]; !ok {
			groups[g] = true
		}
	}

	return len(groups)
}

func addGroup(grid map[int]int, g int) map[int]int {
	gr := g
	if v, ok := grid[g-1]; ok && v < gr {
		gr = v
	}
	if v, ok := grid[g-128]; ok && v < gr {
		gr = v
	}
	grid[g] = gr
	if v, ok := grid[g-1]; ok && v != gr {
		for j, x := range grid {
			if x == v {
				grid[j] = gr
			}
		}
	}
	if v, ok := grid[g-128]; ok && v != gr {
		for j, x := range grid {
			if x == v {
				grid[j] = gr
			}
		}
	}
	return grid
}
