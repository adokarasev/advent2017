package distance

import (
	"math"
)

type point struct {
	x, y float64
}

func (p *point) distance() float64 {
	return math.Abs(p.x) + math.Abs(p.y)
}

// Distance returns Manhattan Distance between
// the location of the data and square 1
func Distance(num int) float64 {
	p := point{0, 0}
	d := point{1, 0}
	matrix := map[point]int{p: 1}
	for i := 1; i < num; i++ {
		p = point{p.x + d.x, p.y + d.y}
		matrix[p] = i
		if _, ok := matrix[point{p.x - d.y, p.y + d.x}]; !ok {
			d = point{-d.y, d.x}
		}
	}
	return p.distance()
}

// Ceil returns first value written that is larger than your puzzle input
func Ceil(num int) int {
	p := point{0, 0}
	d := point{1, 0}
	matrix := map[point]int{p: 1}
	v := 1
	for v < num {
		p = point{p.x + d.x, p.y + d.y}
		v = cellValue(matrix, p)
		matrix[p] = v
		if _, ok := matrix[point{p.x - d.y, p.y + d.x}]; !ok {
			d = point{-d.y, d.x}
		}
	}
	return v
}

func cellValue(matrix map[point]int, p point) int {
	var val int
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if v, ok := matrix[point{p.x + float64(x), p.y + float64(y)}]; ok {
				val += v
			}
		}
	}
	return val
}
