package checksum

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []struct {
	spreadsheet string
	checksum    int
}{
	{
		spreadsheet: `5 1 9 5
		7 5 3
		2 4 6 8`,
		checksum: 18},
}

func TestChecksum(t *testing.T) {
	assert := require.New(t)
	for _, in := range input {
		sum := Checksum(in.spreadsheet)
		assert.Equal(in.checksum, sum)
	}
}

func TestSolveChecksum(t *testing.T) {
	assert := require.New(t)
	dat, _ := ioutil.ReadFile("./input.txt")
	sum := Checksum(string(dat))
	assert.Equal(45972, sum)
}

var edvinput = []struct {
	spreadsheet string
	checksum    int
}{
	{
		spreadsheet: `5 9 2 8
		9 4 7 3
		3 8 6 5`,
		checksum: 9},
}

func TestEvenlyDivisibleValues(t *testing.T) {
	assert := require.New(t)
	for _, in := range edvinput {
		sum := EvenlyDivisibleValues(in.spreadsheet)
		assert.Equal(in.checksum, sum)
	}
}

func TestSolveEvenlyDivisibleValues(t *testing.T) {
	assert := require.New(t)
	dat, _ := ioutil.ReadFile("./input.txt")
	sum := EvenlyDivisibleValues(string(dat))
	assert.Equal(326, sum)
}
