package leetcode

func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	res := make([]int, len(nums))
	ptr1 := 0
	ptr2 := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[ptr1] > nums[ptr2] {
			res[i] = nums[ptr1]
			ptr1++
		} else {
			res[i] = nums[ptr2]
			ptr2--
		}
	}

	return res
}
