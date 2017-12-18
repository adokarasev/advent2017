package defragmentation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountUsedSquares(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"", args{"flqrgnkx"}, 8108},
		{"", args{"stpzcrnm"}, 8250},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			grid := createGrid(tt.args.key)
			assert.Equal(tt.want, CountUsedSquares(grid))
		})
	}
}

func TestCountGroup(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"flqrgnkx"}, 1242},
		{"", args{"stpzcrnm"}, 8250},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			grid := createGrid(tt.args.key)
			assert.Equal(tt.want, CountGroup(grid))
		})
	}
}

func createGrid(key string) []string {
	var grid []string
	for i := 0; i < 128; i++ {
		grid = append(grid, fmt.Sprintf("%s-%d", key, i))
	}
	return grid
}
