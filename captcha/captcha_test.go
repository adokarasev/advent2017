package captcha

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []struct {
	seq string
	sum int
}{
	{"1122", 3},
	{"1111", 4},
	{"1234", 0},
	{"91212129", 9},
}

func TestCaptcha(t *testing.T) {
	assert := require.New(t)

	for _, in := range input {
		sum := Captcha(in.seq, next)
		assert.Equal(in.sum, sum)
	}
}

func TestSolveCaptha(t *testing.T) {
	assert := require.New(t)

	dat, _ := ioutil.ReadFile("./input.txt")
	sum := Captcha(string(dat), next)
	assert.Equal(1393, sum)
}

var hfinput = []struct {
	seq string
	sum int
}{
	{"1212", 6},
	{"1221", 0},
	{"123425", 4},
	{"123123", 12},
	{"12131415", 4},
}

func TestHalfwayCapture(t *testing.T) {
	assert := require.New(t)

	for _, in := range hfinput {
		sum := Captcha(in.seq, halfway)
		assert.Equal(in.sum, sum)
	}
}

func TestSolveHalfwayCaptha(t *testing.T) {
	assert := require.New(t)

	dat, _ := ioutil.ReadFile("./input.txt")
	sum := Captcha(string(dat), halfway)
	assert.Equal(1292, sum)
}
