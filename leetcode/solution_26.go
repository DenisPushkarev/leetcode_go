package leetcode

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ptr := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[ptr] && ptr != i {
			ptr++
			nums[ptr] = nums[i]
		}
	}
	return ptr + 1
}
