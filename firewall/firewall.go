package firewall

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

var digits = regexp.MustCompile(`\d+`)

// Severity returns severity of the whole trip
func Severity(conf string) int {
	var severity int
	layers, count := loadConfig(conf)
	for t := 0; t < count+1; t++ {
		if r, ok := layers[t]; ok {
			if t%(2*(r-1)) == 0 {
				severity += t * r
			}
		}
	}
	return severity
}

func FindDelay(conf string) int {
	layers, _ := loadConfig(conf)
	for d := 0; ; d++ {
		ok := true
		for l, r := range layers {
			if (d+l)%(2*(r-1)) == 0 {
				ok = false
				break
			}
		}

		if ok {
			return d
		}
	}
}

func loadConfig(conf string) (map[int]int, int) {
	var count int
	layers := make(map[int]int)
	lines := bufio.NewScanner(strings.NewReader(conf))

	for lines.Scan() {
		str := lines.Text()
		cfg := digits.FindAllString(str, -1)
		layer, _ := strconv.Atoi(cfg[0])
		rng, _ := strconv.Atoi(cfg[1])
		layers[layer] = rng
		if count < layer {
			count = layer
		}
	}

	return layers, count
}
