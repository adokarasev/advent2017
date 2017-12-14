package stream

// CountScore finds the total score for all groups in your input.
// Each group is assigned a score which is one more than the score of
// the group that immediately contains it. (The outermost group gets a score of 1.)
func CountScore(stream string) (int, int) {
	var garbage, ignore bool
	var totalScore, groupScore, allGarbage int

	for _, r := range stream {
		if ignore {
			ignore = !ignore
			continue
		} else if r == '!' {
			ignore = !ignore
		} else if r == '<' && !garbage {
			garbage = !garbage
		} else if r == '>' {
			garbage = !garbage
		} else if r == '{' && !garbage {
			groupScore++
		} else if r == '}' && !garbage {
			totalScore += groupScore
			groupScore--
		} else if garbage {
			allGarbage++
		}
	}

	return totalScore, allGarbage
}
