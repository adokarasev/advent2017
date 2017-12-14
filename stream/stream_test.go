package stream

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountScore(t *testing.T) {
	type args struct {
		stream string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"{}"}, 1},
		{"", args{"{{{}}}"}, 6},
		{"", args{"{{},{}}"}, 5},
		{"", args{"{{{},{},{{}}}}"}, 16},
		{"", args{"{<a>,<a>,<a>,<a>}"}, 1},
		{"", args{"{{<ab>},{<ab>},{<ab>},{<ab>}}"}, 9},
		{"", args{"{{<!!>},{<!!>},{<!!>},{<!!>}}"}, 9},
		{"", args{"{{<a!>},{<a!>},{<a!>},{<ab>}}"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := CountScore(tt.args.stream); got != tt.want {
				t.Errorf("CountScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountGarbage(t *testing.T) {
	type args struct {
		stream string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"<>"}, 0},
		{"", args{"<random characters>"}, 17},
		{"", args{"<<<<>"}, 3},
		{"", args{"<{!>}>"}, 2},
		{"", args{"<!!>"}, 0},
		{"", args{"<!!!>>"}, 0},
		{"", args{"<{o\"i!a,<{i<a>"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := CountScore(tt.args.stream); got != tt.want {
				t.Errorf("CountScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveCountScore(t *testing.T) {
	assert := require.New(t)
	d, _ := ioutil.ReadFile("./input.txt")
	s, g := CountScore(string(d))
	assert.Equal(12505, s)
	assert.Equal(0, g)
}
