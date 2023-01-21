package leetcode

import (
	"math"
)

func ReverseInt(x int) int {
	return reverseInt(x)
}
func reverseInt(x int) int {
	var pop int
	var result int = 0
	maxVal := math.MaxInt32 / 10
	minVal := math.MinInt32 / 10
	for x != 0 {
		pop = int(x % 10)
		if result > maxVal || result == maxVal && pop > 7 {
			return 0
		}
		if result < minVal || result == minVal && pop < -8 {
			return 0
		}
		result = result*10 + pop
		x = int(math.Round(float64(x / 10)))
	}
	return result
}
