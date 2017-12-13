package reallocation

// ComputeReallocationCycle counts  how many redistribution cycles
// must be completed before a configuration is produced that has been seen before
func ComputeReallocationCycle(banks []int) (int, int) {
	var n int
	cycles := make(map[int]int)
	for n = 1; ; n++ {
		m := findMax(banks)
		blocks := banks[m]
		banks[m] = 0
		for i := 1; i <= blocks; i++ {
			banks[(m+i)%len(banks)]++
		}
		h := hash(banks)
		if v, ok := cycles[h]; ok {
			return n, n - v
		}
		cycles[h] = n
	}
}

func findMax(b []int) int {
	var max int
	for i, v := range b {
		if b[max] < v {
			max = i
		}
	}
	return max
}

func hash(b []int) int {
	var h int
	for _, v := range b {
		h = h*10 + v
	}
	return h
}
