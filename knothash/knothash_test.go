package knothash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKnotHash(t *testing.T) {
	data := []struct {
		alen    int
		lengths []byte
		result  int
	}{
		{5, []byte{3, 4, 1, 5}, 12},
	}

	assert := require.New(t)
	for _, in := range data {
		out := KnotHash(in.alen, in.lengths)
		assert.Equal(in.result, out[0]*out[1])
	}
}

func TestSolveKnotHash(t *testing.T) {
	assert := require.New(t)
	out := KnotHash(256, []byte{106, 118, 236, 1, 130, 0, 235, 254, 59, 205, 2, 87, 129, 25, 255, 118})
	assert.Equal(6909, out[0]*out[1])
}

func TestHash(t *testing.T) {
	data := []struct {
		str  string
		want string
	}{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}

	assert := require.New(t)
	for _, in := range data {
		out := Hash(256, in.str)
		assert.Equal(in.want, out)
	}
}

func TestSolveHash(t *testing.T) {
	assert := require.New(t)
	hash := Hash(256, "106,118,236,1,130,0,235,254,59,205,2,87,129,25,255,118")
	assert.Equal("9d5f4561367d379cfbf04f8c471c0095", hash)
}
