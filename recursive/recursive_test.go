package recursive

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolveBalanceTower(t *testing.T) {
	assert := require.New(t)
	data, _ := ioutil.ReadFile("./input.txt")
	w := BalanceTower(string(data))
	assert.Equal(7, w)
}

func TestBalanceTower(t *testing.T) {
	assert := require.New(t)
	data := `pbga (66)
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
cntj (57)`
	w := BalanceTower(data)
	assert.Equal(8, w)
}
