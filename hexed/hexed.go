package hexed

// Distance determines the fewest number of
// steps required to reach the same point
func Distance(steps []string) int {
	var path []string
	for _, s := range steps {
		path = addStep(path, s)
	}
	return len(path)
}

// MaxDistance returns the furthest distance
func MaxDistance(steps []string) int {
	var max int
	var path []string
	for _, s := range steps {
		path = addStep(path, s)
		if max < len(path) {
			max = len(path)
		}
	}
	return max
}

func addStep(path []string, s string) []string {
	for i, p := range path {
		if r, ok := replaceStep(p, s); ok {
			path = append(path[:i], path[i+1:]...)
			if r != "" {
				return addStep(path, r)
			}
			return path
		}
	}
	return append(path, s)
}

func replaceStep(p string, s string) (string, bool) {
	switch p + s {
	case "ns", "sn", "nesw", "swne", "senw", "nwse":
		return "", true
	case "nenw", "nwne":
		return "n", true
	case "sesw", "swse":
		return "s", true
	case "swn", "nsw":
		return "nw", true
	case "sen", "nse":
		return "ne", true
	case "nws", "snw":
		return "sw", true
	case "nes", "sne":
		return "se", true
	}
	return "", false
}
