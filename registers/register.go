package register

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

const (
	pattern = `([a-z]+) (inc|dec) (-?[0-9]+) if ([a-z]+) ([<>=!]+) (-?[0-9]+)`
)

type condition func(map[string]int, string, int) bool

func gt(registry map[string]int, reg string, val int) bool {
	return registry[reg] > val
}

func lt(registry map[string]int, reg string, val int) bool {
	return registry[reg] < val
}

func eq(registry map[string]int, reg string, val int) bool {
	return registry[reg] == val
}

func ne(registry map[string]int, reg string, val int) bool {
	return registry[reg] != val
}

func ge(registry map[string]int, reg string, val int) bool {
	return registry[reg] >= val
}

func le(registry map[string]int, reg string, val int) bool {
	return registry[reg] <= val
}

var cons = map[string]condition{
	`>`:  gt,
	`<`:  lt,
	`==`: eq,
	`!=`: ne,
	`>=`: ge,
	`<=`: le,
}

type operation func(map[string]int, string, int) int

func inc(registry map[string]int, reg string, val int) int {
	return registry[reg] + val
}

func dec(registry map[string]int, reg string, val int) int {
	return registry[reg] - val
}

var ops = map[string]operation{
	`inc`: inc,
	`dec`: dec,
}

func findMax(instructions string) (int, int) {
	registry := make(map[string]int)
	var highest int
	lines := bufio.NewScanner(strings.NewReader(instructions))
	for lines.Scan() {
		inst := lines.Text()
		if strings.TrimSpace(inst) == "" {
			continue
		}

		re := regexp.MustCompile(pattern)
		tokens := re.FindAllStringSubmatch(inst, -1)
		if len(tokens) > 0 {
			t := tokens[0]
			v, _ := strconv.Atoi(t[6])
			if cons[t[5]](registry, t[4], v) {
				v, _ := strconv.Atoi(t[3])
				h := ops[t[2]](registry, t[1], v)
				if highest < h {
					highest = h
				}
				registry[t[1]] = h
			}
		}
	}

	var max int
	for _, v := range registry {
		if max < v {
			max = v
		}
	}

	return max, highest
}
