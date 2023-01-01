package leetcode

import "fmt"

func RemoveElement(nums []int, val int) int {
	return removeElement(nums, val)
}
func removeElement(nums []int, val int) int {
	i, k := 0, len(nums)-1
	for i <= k {
		if nums[i] == val {
			if nums[k] == val {
				k--
			} else {
				nums[i] = nums[k]
				k--
				i++
			}
		} else {
			i++
		}
		fmt.Println(nums)
	}
	return i
}
