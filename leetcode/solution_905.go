package leetcode

func sortArrayByParity(nums []int) []int {
	j, k := 0, len(nums)-1
	a, b := 0, 0
	for j < k {
		a = nums[j] % 2
		b = nums[k] % 2
		switch {
		case a == 0:
			j++
		case b == 1:
			k--
		case a == 1 && b == 0:
			nums[j], nums[k] = nums[k], nums[j]
		}
	}
	return nums
}
