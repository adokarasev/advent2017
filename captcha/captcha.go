package captcha

// Captcha function returns a string captcha value
func Captcha(seq string, pair func(string, int) rune) int {
	sum := 0
	for i, r := range seq {
		if pair(seq, i) == r {
			sum += int(r - '0')
		}
	}
	return sum
}

func next(seq string, i int) rune {
	return rune(seq[(i+1)%len(seq)])
}

func halfway(seq string, i int) rune {
	len := len(seq)
	return rune(seq[(i+len/2)%len])
}
