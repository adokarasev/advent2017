package recursive

import (
	"bufio"
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

func (t *tower) addChild(c *tower) {
	fmt.Println("Add child", c.name, "to", t.name)
	c.parent = t
	t.children = append(t.children, c)
	t.update(c.weight)
}

// BalanceTower returns the name and weight of program to make tower balanced
func BalanceTower(data string) (string, int) {
	towers := make(map[string]tower)
	lines := bufio.NewScanner(strings.NewReader(data))
	re := regexp.MustCompile(`[a-z0-9]+`)
	for lines.Scan() {
		t := re.FindAllString(lines.Text(), -1)
		w, _ := strconv.Atoi(t[1])

		r, ok := towers[t[0]]
		if !ok {
			fmt.Println("Create new parent", t[0], w)
			r = tower{name: t[0], weight: w}
			towers[t[0]] = r
		} else {
			fmt.Println("Just update it", r)
			r.update(w)
		}

		for _, v := range t[2:] {
			c, ok := towers[v]
			if !ok {
				c = tower{name: v, weight: 0}
				towers[v] = c
			}
			r.addChild(&c)
			fmt.Println("Added child", c, "to", r)
		}
	}

	for _, t := range towers {
		fmt.Println(t)
		if t.parent == nil {
			return t.name, t.weight
		}
	}

	return "", 0
}
