package leetcode

func FindDisappearedNumbers(nums []int) []int {
	return findDisappearedNumbers(nums)
}
func findDisappearedNumbers(nums []int) []int {
	b := make([]int, len(nums)+1)
	l := 0
	for _, val := range nums {
		b[val]++
		if b[val] > 1 {
			l++
		}
	}

	c := make([]int, l)
	ptr := 0
	for i := 1; i < len(b); i++ {
		if b[i] == 0 {
			c[ptr] = i
			ptr++
		}
	}
	return c
}
