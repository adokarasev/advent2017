package tubes

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {
	tubes := `
    |          
    |  +--+    
    A  |  C    
F---|----E|--+ 
    |  |  |  D 
    +B-+  +--+	 
`
	assert := require.New(t)
	path, steps := Path(tubes)
	assert.Equal("ABCDEF", path)
	assert.Equal(38, steps)
}

func TestSolvePath(t *testing.T) {
	data, _ := ioutil.ReadFile("./input.txt")
	assert := require.New(t)
	path, steps := Path(string(data))
	assert.Equal("MKXOIHZNBL", path)
	assert.Equal(17872, steps)
}
