package instruction

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []struct {
	data  []int
	count int
}{
	{[]int{0, 3, 0, 1, -3}, 5},
}

func TestCountSteps(t *testing.T) {
	assert := require.New(t)
	for _, in := range input {
		assert.Equal(in.count, CountSteps(in.data))
	}
}

func loadInput() []int {
	var in []int
	dat, _ := ioutil.ReadFile("./input.txt")
	scanner := bufio.NewScanner(bytes.NewReader(dat))
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		in = append(in, v)
	}
	return in
}

func TestSolveCountSteps(t *testing.T) {
	assert := require.New(t)
	in := loadInput()
	assert.Equal(358309, CountSteps(in))
}

func TestSolveCountStrangeSteps(t *testing.T) {
	assert := require.New(t)
	in := loadInput()
	assert.Equal(28178177, CountStrangeSteps(in))
}
