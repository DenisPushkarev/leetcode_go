package leetcode

import (
	"math"
)

func ThirdMax(nums []int) int {
	return thirdMax(nums)
}
func thirdMax(nums []int) int {
	m1, m2, m3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, val := range nums {
		switch {
		case val > m1:
			m3, m2, m1 = m2, m1, val
		case val > m2 && val != m1:
			m3, m2 = m2, val
		case val > m3 && val != m2 && val != m1:
			m3 = val
		}
	}
	if m3 == math.MinInt64 {
		return m1
	} else {
		return m3
	}
}
