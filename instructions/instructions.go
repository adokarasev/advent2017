package instruction

// CountSteps returns a number of step required to reach the end of slice
// when instruction is incremented by 1 after jump
func CountSteps(instructions []int) int {
	return count(instructions, func(n int) int { return n + 1 })
}

// CountStrangeSteps returns a number of step required to reach the end of slice
// when the jumps are even stranger: after each jump,
// if the offset was three or more, instead decrease it by 1.
// Otherwise, increase it by 1 as before.
func CountStrangeSteps(instructions []int) int {
	return count(instructions, func(n int) int {
		if n >= 3 {
			return n - 1
		}
		return n + 1
	})
}

func count(instructions []int, increment func(int) int) int {
	var n, c = 0, 0
	for i := 0; i < len(instructions); i += n {
		n = instructions[i]
		instructions[i] = increment(instructions[i])
		c++
	}
	return c
}
