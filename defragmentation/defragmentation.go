package defragmentation

import (
	"strconv"

	"github.com/adokarasev/advent2017/knothash"
	"github.com/adokarasev/advent2017/util"
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
	var grid [][]int
	for i, s := range in {
		hash := knothash.Hash(256, s)
		var row []int
		for _, r := range hash {
			j := 1
			n, _ := strconv.ParseInt(string(r), 16, 8)
			for t := 0; t < 4; t++ {
				row = append(row, int(n&1)*((i+1)*j))
				n = n >> 1
				j++
			}
		}
		grid = append(grid, row)
	}

	for i, r := range grid {
		for j, s := range r {
			if s > 0 {
				grid[i][j] = group(grid, i, j)
			}
		}
	}

	groups := make(map[int]bool)
	for _, r := range grid {
		for _, s := range r {
			if s > 0 {
				groups[s] = true
			}
		}
	}

	return len(grid)
}

func group(grid [][]int, i int, j int) int {
	g := grid[i][j]
	for x := util.MaxInt(0, i-1); x < util.MinInt(len(grid), i+2); x++ {
		for y := util.MaxInt(0, j-1); y < util.MinInt(len(grid[i]), j+2); y++ {
			if grid[x][y] > 0 && g > grid[x][y] {
				g = grid[x][y]
			}
		}
	}
	return g
}
