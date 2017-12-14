package register

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMax(t *testing.T) {
	assert := require.New(t)
	m, h := findMax(`
		b inc 5 if a > 1
		a inc 1 if b < 5
		c dec -10 if a >= 1
		c inc -20 if c == 10`)

	assert.Equal(1, m)
	assert.Equal(10, h)
}

func TestSolveFindMax(t *testing.T) {
	assert := require.New(t)
	data, _ := ioutil.ReadFile("./input.txt")
	m, h := findMax(string(data))
	assert.Equal(6061, m)
	assert.Equal(0, h)
}
