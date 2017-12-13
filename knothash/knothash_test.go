package knothash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKnotHash(t *testing.T) {
	data := []struct {
		alen    int
		lengths []int
		result  int
	}{
		{5, []int{3, 4, 1, 5}, 12},
	}

	assert := require.New(t)
	for _, in := range data {
		out := KnotHash(in.alen, in.lengths)
		assert.Equal(in.result, out[0]*out[1])
	}
}

func TestSolveKnotHash(t *testing.T) {
	assert := require.New(t)
	out := KnotHash(256, []int{106, 118, 236, 1, 130, 0, 235, 254, 59, 205, 2, 87, 129, 25, 255, 118})
	assert.Equal(6909, out[0]*out[1])
}
