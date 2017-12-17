package plumber

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

var digits = regexp.MustCompile(`\d+`)

// FindConnections returns a number of elements connected to given node
func FindConnections(data string, node int) int {
	network := buildNetwork(data)
	if net, ok := network[node]; ok {
		var counter int
		for _, v := range network {
			if net == v {
				counter++
			}
		}
		return counter
	}
	return 0
}

// CountGroups returns number of groups in the network
func CountGroups(data string) int {
	network := buildNetwork(data)
	groups := make(map[int]bool)
	for _, net := range network {
		if _, ok := groups[net]; !ok {
			groups[net] = true
		}
	}
	return len(groups)
}

func buildNetwork(data string) map[int]int {
	lines := bufio.NewScanner(strings.NewReader(data))
	network := make(map[int]int)
	for lines.Scan() {
		t := lines.Text()
		var net int
		for i, s := range digits.FindAllString(t, -1) {
			n, _ := strconv.Atoi(s)
			if i == 0 {
				if _, ok := network[n]; !ok {
					network[n] = n
					net = n
				} else {
					net, _ = network[n]
				}
			} else {
				if net1, ok := network[n]; !ok {
					network[n] = net
				} else {
					for k, v := range network {
						if net1 == v {
							network[k] = net
						}
					}
				}
			}
		}
	}
	return network
}
