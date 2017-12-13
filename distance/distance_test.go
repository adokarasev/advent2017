package distance

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []struct {
	point int
	dist  float64
}{
	{1, 0},
	{12, 3},
	{23, 2},
	{1024, 31},
	{312051, 430},
}

func TestDistance(t *testing.T) {
	assert := require.New(t)
	for _, in := range input {
		assert.Equal(in.dist, Distance(in.point))
	}
}

func TestCeil(t *testing.T) {
	assert := require.New(t)
	assert.Equal(312453, Ceil(312051))
}
