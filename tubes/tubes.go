package tubes

import (
	"bufio"
	"errors"
	"strings"
)

type position struct {
	x, y int
}

type direction struct {
	x, y int
}

func (p *position) add(d direction) *position {
	return &position{p.x + d.x, p.y + d.y}
}

// Path returns letters it will see (in the order it would see them) if it follows the path
// and number of steps
func Path(tubesMap string) (string, int) {
	scanner := bufio.NewScanner(strings.NewReader(tubesMap))
	var tMap [][]rune
	for scanner.Scan() {
		txt := scanner.Text()
		if len(strings.TrimSpace(txt)) > 0 {
			tMap = append(tMap, []rune(txt))
		}
	}

	var path []rune
	var steps int
	d := direction{0, 1}
	for p, _ := findEnterance(tMap); ; p = p.add(d) {
		if r, err := get(tMap, p); err == nil {
			switch r {
			case '|', '-':
			case '+':
				d = direction{-d.y, d.x}
				n := p.add(d)
				if _, err = get(tMap, n); err != nil {
					d = direction{-d.x, -d.y}
				}
			default:
				path = append(path, r)
			}
			steps++
		} else {
			break
		}
	}

	return string(path), steps
}

func findEnterance(tMap [][]rune) (*position, error) {
	for x, r := range tMap[0] {
		if r == '|' {
			return &position{x, 0}, nil
		}
	}
	return nil, errors.New("Can't find enterance")
}

func get(tMap [][]rune, p *position) (rune, error) {
	if p.y >= 0 && p.y < len(tMap) && p.x >= 0 && p.x < len(tMap[p.y]) && tMap[p.y][p.x] != ' ' {
		return tMap[p.y][p.x], nil
	}
	return ' ', errors.New("No path element in this position")
}
