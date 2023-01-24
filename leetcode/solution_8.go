package leetcode

import (
	"math"
)

func myAtoi(s string) int {
	var isNegative byte
	var result int
	var tmp int
	var ptr int = 0
	for ; ptr < len(s) && s[ptr] == ' '; ptr++ {
	}
	if ptr < len(s) {
		if s[ptr] == '-' {
			isNegative = 1
			ptr++
		} else if s[ptr] == '+' {
			ptr++
		}
	}
	for ; ptr < len(s) && s[ptr] == '0'; ptr++ {
	}
	for i := ptr; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			tmp = int(s[i] - '0')
			if result > 214748364 {
				return max(isNegative)
			}
			result *= 10
			if isNegative == 1 && math.MinInt32+tmp > result {
				return max(isNegative)
			} else if math.MaxInt32-tmp < result {
				return max(isNegative)
			} else {
				result += tmp
			}
		} else {
			return withSign(result, isNegative)
		}
	}
	return withSign(result, isNegative)
}
func withSign(number int, isNegative byte) int {
	if isNegative == 1 {
		return -number
	} else {
		return number
	}
}

func max(isNegative byte) int {
	if isNegative == 1 {
		return math.MinInt32
	} else {
		return math.MaxInt32
	}
}
