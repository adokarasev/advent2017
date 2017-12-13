package reallocation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var data = []struct {
	banks []int
	index int
	size  int
}{
	{[]int{0, 2, 7, 0}, 5, 4},
	{[]int{4, 1, 15, 12, 0, 9, 9, 5, 5, 8, 7, 3, 14, 5, 12, 3}, 6681, 0},
}

func TestComputeReallocationCycle(t *testing.T) {
	assert := require.New(t)
	for _, in := range data {
		n, s := ComputeReallocationCycle(in.banks)
		assert.Equal(in.index, n)
		assert.Equal(in.size, s)
	}
}
