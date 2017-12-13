package recursive

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = struct {
	data string
	root string
}{
	`pbga (66)
	xhth (57)
	ebii (61)
	havc (66)
	ktlj (57)
	fwft (72) -> ktlj, cntj, xhth
	qoyq (66)
	padx (45) -> pbga, havc, qoyq
	tknk (41) -> ugml, padx, fwft
	jptl (61)
	ugml (68) -> gyxo, ebii, jptl
	gyxo (61)
	cntj (57)`,
	"tknk"}

func TestFindTowerRoot(t *testing.T) {
	assert := require.New(t)
	assert.Equal(input.root, FindRoot(input.data))
}

func TestSolveFindTowerRoot(t *testing.T) {
	assert := require.New(t)
	data, _ := ioutil.ReadFile("./input.txt")
	assert.Equal("veboyvy", FindRoot(string(data)))
}

func TestBalanceTower(t *testing.T) {
	assert := require.New(t)
	assert.Equal(input.root, FindRoot(input.data))
}

func TestSolveBalanceTower(t *testing.T) {
	assert := require.New(t)
	data, _ := ioutil.ReadFile("./input.txt")
	s, w := BalanceTower(string(data))
	assert.Equal("veboyvy", s)
	assert.Equal(1, w)
}
