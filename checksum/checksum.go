package checksum

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

// Checksum returns a spreadsheet checksum.
// For each row, determine the difference between the largest
// value and the smallest value; the checksum is the sum of all
// of these differences
func Checksum(spreadsheet string) int {
	var checksum = 0
	var hi = 0
	var lo = math.MaxInt32

	var calculator = func(n int) {
		if hi < n {
			hi = n
		}
		if lo > n {
			lo = n
		}
	}

	calculate(spreadsheet, calculator, func() {
		checksum += hi - lo
		hi = 0
		lo = math.MaxInt32
	})
	return checksum
}

// EvenlyDivisibleValues finds the only two numbers in each row
// where one evenly divides the other - that is,
// where the result of the division operation is a whole number.
// They would like you to find those numbers on each line, divide them,
// and add up each line's result
func EvenlyDivisibleValues(spreadsheet string) int {
	var checksum = 0
	var nums []int
	calculate(spreadsheet, func(n int) { nums = append(nums, n) }, func() {
		for i, a := range nums {
			for _, b := range nums[i+1:] {
				if a%b == 0 || b%a == 0 {
					checksum += a/b + b/a
					nums = nil
					return
				}
			}
		}
	})
	return checksum
}

func calculate(spreadsheet string, f func(n int), accumulator func()) {
	scanner := bufio.NewScanner(strings.NewReader(spreadsheet))
	for scanner.Scan() {
		nums := bufio.NewScanner(strings.NewReader(scanner.Text()))
		nums.Split(bufio.ScanWords)
		for nums.Scan() {
			n, _ := strconv.Atoi(nums.Text())
			f(n)
		}
		accumulator()
	}
}
