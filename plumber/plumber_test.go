package plumber

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

var data = `
0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5
`

func TestFindConnections(t *testing.T) {
	assert := require.New(t)
	assert.Equal(6, FindConnections(data, 0))
}

func TestSolveFindConnections(t *testing.T) {
	assert := require.New(t)
	data, _ := ioutil.ReadFile("./input.txt")
	assert.Equal(152, FindConnections(string(data), 0))
}

func TestCountGroups(t *testing.T) {
	assert := require.New(t)
	assert.Equal(2, CountGroups(data))
}

func TestSolveCountGroups(t *testing.T) {
	assert := require.New(t)
	data, _ := ioutil.ReadFile("./input.txt")
	assert.Equal(186, CountGroups(string(data)))
}
