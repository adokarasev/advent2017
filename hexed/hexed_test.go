package hexed

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDistance(t *testing.T) {
	type args struct {
		steps []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"ne", "ne", "ne"}}, 3},
		{"2", args{[]string{"ne", "ne", "sw", "sw"}}, 0},
		{"3", args{[]string{"ne", "ne", "s", "s"}}, 2},
		{"4", args{[]string{"se", "sw", "se", "sw", "sw"}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			if got := Distance(tt.args.steps); got != tt.want {
				assert.Equal(tt.want, got)
			}
		})
	}
}

func TestSolveDistance(t *testing.T) {
	assert := require.New(t)
	data, _ := ioutil.ReadFile("./input.txt")
	d := Distance(strings.Split(string(data), ","))
	assert.Equal(810, d)
}

func TestSolveMaxDistance(t *testing.T) {
	assert := require.New(t)
	data, _ := ioutil.ReadFile("./input.txt")
	d := MaxDistance(strings.Split(string(data), ","))
	assert.Equal(1567, d)
}
