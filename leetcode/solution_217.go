package leetcode

import "sort"

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sort.Ints(nums)
	prev := nums[0]
	uniq := false
	for i := 1; i < len(nums) && !uniq; i++ {
		if nums[i] == prev {
			uniq = true
		}
		prev = nums[i]
	}
	return uniq
}
