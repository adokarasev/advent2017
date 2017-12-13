package passphrase

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []struct {
	in    string
	valid bool
}{
	{"aa bb cc dd ee", true},
	{"aa bb cc dd aa", false},
	{"aa bb cc dd aaa", true},
}

func TestPassphase(t *testing.T) {
	assert := require.New(t)
	for _, in := range input {
		assert.Equal(in.valid, Passphrase(in.in))
	}
}

func TestSolvePassphrase(t *testing.T) {
	assert := require.New(t)
	dat, _ := ioutil.ReadFile("./input.txt")
	n := 0
	scanner := bufio.NewScanner(bytes.NewReader(dat))
	for scanner.Scan() {
		if Passphrase(scanner.Text()) {
			n++
		}
	}
	assert.Equal(386, n)
}

var stinput = []struct {
	in    string
	valid bool
}{
	{"abcde fghij", true},
	{"abcde xyz ecdab ", false},
	{"a ab abc abd abf abj", true},
	{"iiii oiii ooii oooi oooo", true},
	{"oiii ioii iioi iiio ", false},
}

func TestStrongPassphase(t *testing.T) {
	assert := require.New(t)
	for _, in := range stinput {
		assert.Equal(in.valid, StrongPasshrase(in.in))
	}
}

func TestSolveStrongPasshrase(t *testing.T) {
	assert := require.New(t)
	dat, _ := ioutil.ReadFile("./input.txt")
	n := 0
	scanner := bufio.NewScanner(bytes.NewReader(dat))
	for scanner.Scan() {
		if StrongPasshrase(scanner.Text()) {
			n++
		}
	}
	assert.Equal(208, n)
}
