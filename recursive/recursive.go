package recursive

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// FindRoot returns the name of the root program
func FindRoot(data string) string {
	tower := make(map[string]bool)
	lines := bufio.NewScanner(strings.NewReader(data))
	re := regexp.MustCompile(`[a-z]+`)
	for lines.Scan() {
		t := re.FindAllString(lines.Text(), -1)
		v, ok := tower[t[0]]
		tower[t[0]] = !ok || (v && true)
		for _, n := range t[1:] {
			tower[n] = false
		}
	}

	for n, v := range tower {
		if v {
			return n
		}
	}

	return ""
}

type tower struct {
	name     string
	weight   int
	parent   *tower
	children []*tower
}

func (t *tower) update(w int) {
	t.weight += w
	if t.parent != nil {
		t.parent.update(w)
	}
}

func (t *tower) addChild(c *tower) *tower {
	c.parent = t
	t.children = append(t.children, c)
	t.update(c.weight)
	return c
}

func (t *tower) String() string {
	return fmt.Sprint("{", t.name, ",", t.weight, ",", len(t.children), "}")
}

// BalanceTower returns the name and weight of program to make tower balanced
func BalanceTower(data string) int {
	towers := make(map[string]*tower)
	lines := bufio.NewScanner(strings.NewReader(data))
	re := regexp.MustCompile(`[a-z0-9]+`)
	for lines.Scan() {
		t := re.FindAllString(lines.Text(), -1)
		w, _ := strconv.Atoi(t[1])
		r, ok := towers[t[0]]
		if !ok {
			r = &tower{name: t[0], weight: w}
		} else {
			r.update(w)
		}

		for _, v := range t[2:] {
			c, ok := towers[v]
			if !ok {
				c = &tower{name: v, weight: 0}
			}
			towers[v] = r.addChild(c)
		}

		towers[t[0]] = r
	}

	r, _ := getTowerRoot(towers)
	var (
		minW = r.children[0].weight
		maxW = r.children[0].weight
	)

	for _, n := range r.children {
		if minW > n.weight {
			minW = n.weight
		}
		if maxW < n.weight {
			maxW = n.weight
		}
	}

	return maxW - minW
}

func getTowerRoot(towers map[string]*tower) (*tower, error) {
	for _, t := range towers {
		if t.parent == nil {
			return t, nil
		}
	}

	return nil, errors.New("No root found")
}
