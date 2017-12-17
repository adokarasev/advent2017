package firewall

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

var data = `0: 3
1: 2
4: 4
6: 4
`

func TestSeverity(t *testing.T) {
	assert := require.New(t)
	assert.Equal(24, Severity(data))
}

func TestSolveSeverity(t *testing.T) {
	assert := require.New(t)
	in, _ := ioutil.ReadFile("./input.txt")
	assert.Equal(1640, Severity(string(in)))
}

func TestFindDelay(t *testing.T) {
	assert := require.New(t)
	assert.Equal(10, FindDelay(data))
}

func TestSolveFindDelay(t *testing.T) {
	assert := require.New(t)
	in, _ := ioutil.ReadFile("./input.txt")
	assert.Equal(3960702, FindDelay(string(in)))
}
