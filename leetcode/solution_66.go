package leetcode

func PlusOne(digits []int) []int {
	return plusOne(digits)
}
func plusOne(digits []int) []int {
	var over int = 1
	for i := len(digits) - 1; i >= 0 && over > 0; i-- {
		digits[i] += over
		if digits[i] > 9 {
			digits[i] = 10 - digits[i]
			over = 1
		} else {
			over = 0
		}
	}
	if over > 0 {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}
